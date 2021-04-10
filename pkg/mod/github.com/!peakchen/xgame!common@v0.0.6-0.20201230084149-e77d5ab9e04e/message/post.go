package message

// add by stefan

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"github.com/Peakchen/xgameCommon/akLog"
	"tool"
	"io/ioutil"
	"net/http"
)

func HttpPostJsonMsg(url string, src interface{}) {
	data, err := json.Marshal(src)
	if err != nil {
		akLog.Error("json marshal data fail, err: ", err)
		return
	}

	buf := bytes.NewBuffer(data)
	_, err = http.Post(url, jsonPostType, buf)
	if err != nil {
		akLog.Error("http post fail, err: ", err)
		return
	}
}

func HttpPostXMLMsg(url string, src interface{}) {
	data, err := json.Marshal(src)
	if err != nil {
		akLog.Error("json marshal data fail, err: ", err)
		return
	}

	buf := bytes.NewBuffer(data)
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, buf)
	if err != nil {
		akLog.Error("NewRequest data fail, err: ", err)
		return
	}

	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", "设置cookie")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	resp, err := client.Do(reqest)
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		akLog.Error("cookie:", cookie)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		akLog.Error("ReadAll data fail, err: ", err)
		return
	}

}

func HttpResponse(w http.ResponseWriter, headcontent []*THttpResponseHead, statusCode int, src string) {
	akLog.FmtPrintln("response: ", statusCode, len([]rune(src)))
	// data, err := json.Marshal(src)
	// if err != nil {
	// 	akLog.Error("json marshal data fail, err: ", err)
	// 	return
	// }

	for _, headitem := range headcontent {
		w.Header().Set(headitem.HeadKey, headitem.HeadValue)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(src))
	if err != nil {
		akLog.Error("write post fail.")
	}
}
