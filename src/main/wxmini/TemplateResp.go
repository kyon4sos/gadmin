package wxmini

/**
订阅消息模板返回结构体
errcode		number	错误码
errmsg		string	错误信息
priTmplId	string	添加至帐号下的模板id，发送小程序订阅消息时所需
 */
type SubTemplate struct {
	*BaseResp
	PriTmplId string `json:"priTmplId"`
}
/**
属性		类型	说明
tid	number	模版标题 id
title	string	模版标题
type	number	模版类型，2 为一次性订阅，3 为长期订阅
categoryId	number	模版所属类目 id
 */
type SubTemplateTitle struct {
	Tid          int    `json:"tid"`
	Title        string `json:"title"`
	TemplateType int    `json:"type"`
	CategoryId   int `json:"categoryId"`
}
