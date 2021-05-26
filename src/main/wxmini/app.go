package wxmini

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

/**
属性			类型	默认值	必填		说明
appid		string		是		小程序 appId
secret		string		是		小程序 appSecret
js_code		string		是		登录时获取的 code
grant_type	string		是		授权类型，此处只需填写 authorization_code
*/
type App struct {
	AppId     string `json:"app_id"`
	Secret    string `json:"secret"`
	AppKey    string `json:"app_key"`
	FreshTime time.Time `json:"fresh_time"`
	client *resty.Client
	*AccessToken
}

var appContainer = map[string]*App{}

func (app *App) Register(key string) {
	if appContainer == nil {
		appContainer = map[string]*App{}
	}
	appContainer[key] = app
}
func NewApp(appid string, secret string) (*App, error) {
	client := resty.New().SetHostURL(wxHost)
	app := &App{AppId: appid, Secret: secret, FreshTime: time.Now(),client: client}
	res, err := GetAccessToken(app)
	if err != nil {
		return nil, err
	}
	app.AccessToken = res
	go NewSchedule("@every 1h", app.autoFreshToken)
	return app, nil
}
func GetApp(key string) *App {
	return appContainer[key]
}

const remain = 160

func (app *App) GetAccessToken() (*AccessToken, error) {
	var exp time.Duration
	if app.ExpiresIn > remain {
		exp = time.Duration(app.ExpiresIn - remain)
	} else {
		exp = time.Duration(7500 - remain)
	}
	expired := app.FreshTime.Add(exp * time.Second)
	if time.Now().Before(expired) {
		return app.AccessToken, nil
	}
	token, err := GetAccessToken(app)
	if err != nil {
		return nil, err
	}
	app.AccessToken = token
	return token, nil
}

func (app *App) GetUnlimitedWxaCode(data *QrcodeReq) (interface{}, error) {

	res, err := GetUnlimitedWxaCodeByte(app,data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (app *App) autoFreshToken() {
	token, err := app.GetAccessToken()
	if token != nil {
		fmt.Printf("刷新access_token %s\n", token.Token)
	}
	if err != nil {
		fmt.Printf("Get Access Token err %s\n", err.Error())
	}
	app.AccessToken = token
	app.FreshTime = time.Now()
	fmt.Printf("app %v\n", app)
}

func (app *App) SubMessageSend(data *SubMessageReq) (interface{}, error) {
	res, err := SubMessageSend(app,data)
	return res,err
}