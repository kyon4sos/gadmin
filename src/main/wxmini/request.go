package wxmini

/**
二维码请求参数
属性	类型	默认值	必填	说明
access_token	string		是	接口调用凭证
scene	string		是	最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
page	string	主页	否	必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
width	number	430	否	二维码的宽度，单位 px，最小 280px，最大 1280px
auto_color	boolean	false	否	自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
line_color	Object	{"r":0,"g":0,"b":0}	否	auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
is_hyaline	boolean	false	否	是否需要透明底色，为 true 时，生成透明底色的小程序
 */
type QrcodeReq struct {
	Scene       string `json:"scene"`
	Page        string `json:"page"`
	Width       int `json:"width"`
	AutoColor   bool `json:"auto_color"`
	LineColor   interface{} `json:"line_color"`
	IsHyaline   bool `json:"is_hyaline"`
}
/**
新增订阅消息模板请求参数
属性	类型			默认值			必填		说明
access_token	string			是		接口调用凭证
tid				string			是		模板标题 id，可通过接口获取，也可登录小程序后台查看获取
kidList			Array.<number>	是		开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
sceneDesc		string			否		服务场景描述，15个字以内
 */
type TemplateReq struct {
	Tid       string `json:"tid"`
	KidList   string `json:"kidList"`
	SceneDesc string `json:"sceneDesc"`
}

/**
subscribeMessageSend.getPubTemplateTitleList
获取帐号所属类目下的公共模板标题
请求参数
属性	类型	默认值	必填	说明
access_token	string		是	接口调用凭证
ids	string		是	类目 id，多个用逗号隔开
start	number		是	用于分页，表示从 start 开始。从 0 开始计数。
limit	number		是	用于分页，表示拉取 limit 条记录。最大为 30。
 */
type TemplateTitleListReq struct {
	Ids   string `json:"ids"`
	Start string   `json:"start"`
	Limit string    `json:"limit"`
}