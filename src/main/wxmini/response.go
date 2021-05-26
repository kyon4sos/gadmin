package wxmini

/**
基本响应类型
返回的 JSON 数据包
属性	类型	说明
errcode	number	错误码
errmsg	string	错误信息
priTmplId	string	添加至帐号下的模板id，发送小程序订阅消息时所需
 */
type BaseResp struct {
	Errcode int `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}


type SubMsgResp struct {
	BaseResp
}
