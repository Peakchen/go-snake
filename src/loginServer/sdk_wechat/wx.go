package sdk_wechat

import (
	"encoding/json"
	"fmt"
	"go-snake/common"
	"go-snake/common/akOrm"
	"go-snake/common/wechat"
	"go-snake/loginServer/base"
	"go-snake/loginServer/entityBase"
	"go-snake/loginServer/sdk_wechat/wechat_model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Peakchen/xgameCommon/akLog"
)

type (
	wxApp struct {
		Addr      string
		AppID     string
		AppSecret string
	}
)

var wx *wxApp

func Run(addr, appid, appSecret string) {
	wx = &wxApp{
		Addr:      addr,
		AppID:     appid,
		AppSecret: appSecret,
	}
	http.HandleFunc("/wxCodeSession", wxCodeSessionRequest)
	http.HandleFunc("/wxAccessToken", wxGetAccessToken)
	http.HandleFunc("/wxCheckSessionKey", wxCheckSessionKey)
	go func() {
		http.ListenAndServe(addr, nil)
	}()
}

func getUrlData(requrl string) []string {
	requrls := strings.Split(requrl, "?")
	if len(requrls) <= 1 {
		akLog.Error("invalid wxCodeSession content, url: ", requrl)
		return nil
	}
	return requrls
}

/*
	step 1
	get sessionkey,openid...
*/
func wxCodeSessionRequest(w http.ResponseWriter, r *http.Request) {
	requrls := getUrlData(r.RequestURI)
	if len(requrls) <= 1 {
		akLog.Error("invalid wxCodeSession content, url: ", r.RequestURI)
		return
	}
	strcode := strings.Split(requrls[1], "=")
	if len(strcode) <= 1 {
		akLog.Error("invalid wxCodeSession content, code: ", requrls[1])
		return
	}
	if strcode[1] == "undefined" {
		return
	}
	url := fmt.Sprintf(wechat.CodeSessionUrl, wx.AppID, wx.AppSecret, strcode[1])
	rsp, err := http.Get(url)
	if err != nil {
		akLog.Error("code session request fail, info: ", url, err)
		return
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		akLog.Error("read resp fail, info: ", rsp.Body, err)
		return
	}

	var content = &wechat.RespCode2Session{}
	err = json.Unmarshal(body, content)
	if err != nil {
		akLog.Error("json unmarshal fail, info: ", body, err)
		return
	}
	if content.Errcode != 0 {
		akLog.Error(url, content.Errcode, content.Errmsg)
	}
	w.Write(body)
}

/*
	step 2
	get access token
*/
func wxGetAccessToken(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(wechat.AccessTokenUrl, wx.AppID, wx.AppSecret)
	rsp, err := http.Get(url)
	if err != nil {
		akLog.Error("code session request fail, info: ", url, err)
		return
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		akLog.Error("read resp fail, info: ", rsp.Body, err)
		return
	}
	var content = &wechat.RespAccessToken{}
	err = json.Unmarshal(body, content)
	if err != nil {
		akLog.Error("json unmarshal fail, info: ", body, err)
		return
	}
	if content.Errcode != 0 {
		akLog.Error(url, content.Errcode, content.Errmsg)
	}
	w.Write(body)
}

/*
	step 3
	check token,save user info to db
*/
func wxCheckSessionKey(w http.ResponseWriter, r *http.Request) {
	akLog.Info("wxCheckSessionKey data, info: ", r.RequestURI)
	sreq, err := url.QueryUnescape(r.RequestURI)
	if err != nil {
		akLog.Error("invalid err: ", err)
		return
	}
	requrls := getUrlData(sreq)
	if len(requrls) <= 1 {
		akLog.Error("invalid CheckSessionKey content, url: ", sreq)
		return
	}
	akLog.Info("wxCheckSessionKey data, info: ", requrls)
	strdata := strings.Split(requrls[1], "&")
	if len(strdata) <= 1 {
		akLog.Error("invalid CheckSessionKey content, code: ", requrls[1])
		return
	}
	if len(strdata) < 4 {
		akLog.Error("invalid requrls, info: ", requrls)
		return
	}
	userinfo := strings.Split(strdata[0], "=")
	if len(userinfo) <= 1 {
		akLog.Error("invalid userinfo, info: ", userinfo)
		return
	}
	if userinfo[1] == "undefined" {
		akLog.Error("invalid userinfo, info: ", userinfo)
		return
	}
	sessionKeySrc := strings.Split(strdata[1], "=")
	if len(sessionKeySrc) <= 1 {
		akLog.Error("invalid sessionKey, info: ", sessionKeySrc)
		return
	}

	sessionKeyIdx := strings.IndexAny(strdata[1], "=")
	sessionKey := strdata[1][sessionKeyIdx+1:]
	if sessionKey == "undefined" {
		akLog.Error("invalid sessionKey, info: ", sessionKeySrc)
		return
	}

	tokenSrc := strings.Split(strdata[2], "=")
	if len(tokenSrc) <= 1 {
		akLog.Error("invalid token, info: ", tokenSrc)
		return
	}
	tokenIdx := strings.IndexAny(strdata[2], "=")
	token := strdata[2][tokenIdx+1:]
	if token == "undefined" {
		akLog.Error("invalid token, info: ", token)
		return
	}

	openidsrc := strings.Split(strdata[3], "=")
	if len(openidsrc) <= 1 {
		akLog.Error("invalid openid, info: ", openidsrc)
		return
	}
	openidIdx := strings.IndexAny(strdata[3], "=")
	openid := strdata[3][openidIdx+1:]
	if openid == "undefined" {
		akLog.Error("invalid openid, info: ", openid)
		return
	}

	signaturesrc := strings.Split(strdata[4], "=")
	if len(signaturesrc) <= 1 {
		akLog.Error("invalid signature, info: ", signaturesrc)
		return
	}
	signatureIdx := strings.IndexAny(strdata[4], "=")
	signature := strdata[4][signatureIdx+1:]
	if signature == "undefined" {
		akLog.Error("invalid signature, info: ", signature)
		return
	}

	encryptedDatasrc := strings.Split(strdata[5], "=")
	if len(encryptedDatasrc) <= 1 {
		akLog.Error("invalid encryptedData, info: ", encryptedDatasrc)
		return
	}
	encryptedDataIdx := strings.IndexAny(strdata[5], "=")
	encryptedData := strdata[5][encryptedDataIdx+1:]
	if encryptedData == "undefined" {
		akLog.Error("invalid encryptedData, info: ", encryptedData)
		return
	}

	ivsrc := strings.Split(strdata[6], "=")
	if len(ivsrc) <= 1 {
		akLog.Error("invalid iv, info: ", ivsrc)
		return
	}
	ivIdx := strings.IndexAny(strdata[6], "=")
	iv := strdata[6][ivIdx+1:]
	if iv == "undefined" {
		akLog.Error("invalid iv, info: ", iv)
		return
	}

	sig := common.HmacSha256("", sessionKey)
	url := fmt.Sprintf(wechat.CheckSessionKeyUrl, token, sig, openid)
	rsp, err := http.Get(url)
	if err != nil {
		akLog.Error("code session request fail, info: ", url, err)
		return
	}
	defer rsp.Body.Close()
	akLog.Info("url:", url)
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		akLog.Error("read resp fail, info: ", rsp.Body, err)
		return
	}
	var content = &wechat.CommonErrorCode{}
	err = json.Unmarshal(body, content)
	if err != nil {
		akLog.Error("json unmarshal fail, info: ", body, err)
		return
	}
	if content.Errcode != 0 {
		akLog.Error(url, content.Errcode, content.Errmsg)
	} else {
		wxLogin(sessionKey, encryptedData, iv)
	}

	w.Write(body)
}

func wxLogin(sessionkey string, encryptedData string, iv string) {
	decryptData, err := wechat.WxDecrypt(sessionkey, encryptedData, iv)
	if err != nil {
		akLog.Error("WxDecrypt fail, err: ", err)
		return
	}
	var userInfo wechat.WxUserInfo
	err = json.Unmarshal(decryptData, &userInfo)
	if err != nil {
		akLog.Error("Unmarshal fail, err: ", err)
		return
	}
	if userInfo.Watermark.AppID != wx.AppID {
		akLog.Error(wechat.ErrAppIDNotMatch)
		return
	}
	//save userinfo to db.
	exist, err := akOrm.HasExistForWx(&wechat_model.WxRole{}, userInfo.OpenID)
	if err != nil {
		akLog.Error("err: ", err)
		return
	}
	if !exist {
		wxrole := wechat_model.NewRoleAtWX(&userInfo)
		if wxrole == nil {
			akLog.Error("new wx role fail, openid: ", userInfo.OpenID)
			return
		}
		entity := entityBase.InitEntity(wxrole.GetDBID())
		entity.LoadWxRole(wxrole)
		base.AddUser(wxrole.GetDBID(), entity)
	} else {
		var wxRole = &wechat_model.WxRole{}
		err := akOrm.GetModel(wxRole, userInfo.OpenID)
		if err != nil {
			akLog.Error("err: ", err)
			return
		}
		wxRole.Copy(&userInfo)
		wxRole.Update()

		role := base.GetUserByID(wxRole.GetDBID())
		if role == nil {
			akLog.Error("get entity fail, dbid: ", wxRole.GetDBID())
			return
		}
		role.LoadWxRole(wxRole)
	}
}
