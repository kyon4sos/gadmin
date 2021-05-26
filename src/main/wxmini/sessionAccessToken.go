package wxmini
/**
属性				类型		说明
access_token	string	获取到的凭证
expires_in		number	凭证有效时间，单位：秒。目前是7200秒之内的值。
errcode			number	错误码
errmsg			string	错误信息
*/
type SessionOpenid struct {
	SessionKey string `json:"session_key"`
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
}

type AccessToken struct {
	ExpiresIn int    `json:"expires_in"`
	Token     string `json:"access_token"`
}
