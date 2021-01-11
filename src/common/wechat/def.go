package wechat

type CommonErrorCode struct {
	Errcode int    `json:"errcode"` //错误码
	Errmsg  string `json:"errmsg"`  //错误信息
}

//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.code2Session.html
/*
	openid		string	用户唯一标识
	session_key	string	会话密钥
	unionid		string	用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	errcode		number	错误码
	errmsg		string	错误信息
*/
type RespCode2Session struct {
	CommonErrorCode

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回

}

/*
	access_token	string	获取到的凭证
	expires_in		number	凭证有效时间，单位：秒。目前是7200秒之内的值。
	errcode			number	错误码
	errmsg			string	错误信息
*/
type RespAccessToken struct {
	CommonErrorCode

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.code2Session.html
	CodeSessionUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"

	//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/login/auth.checkSessionKey.html
	CheckSessionKeyUrl = "https://api.weixin.qq.com/wxa/checksession?access_token=%v&signature=%v&openid=%v&sig_method=hmac_sha256"

	//https://developers.weixin.qq.com/minigame/dev/api-backend/open-api/access-token/auth.getAccessToken.html
	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
)

var (
	AppID     = "wx75facc388e778570"
	AppSecret = "5cd5ce4cb612f4bb24e1e317b425128c"
)

type WxUserInfo struct {
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}
