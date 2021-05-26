package wxmini

/**
请求参数
属性	类型	默认值	必填	说明
access_token	string		是	接口调用凭证
touser	string		是	接收者（用户）的 openid
template_id	string		是	所需下发的订阅模板id
page	string		否	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
data	Object		是	模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
miniprogram_state	string		否	跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
lang	string		否	进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
 */
type SubMessageReq struct {
	Touser            string            `json:"touser"`
	TemplateId        string            `json:"template_id"`
	Page              string            `json:"page"`
	Data              map[string]map[string]interface{} `json:"data"`
	Miniprogram_state string            `json:"miniprogram_state"`
	Lang              string            `json:"lang"`
}

type MessageData struct {
	Key map[string]string
}

