package wxmini

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"time"
)

/**
GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
*/
func GetCode2Session(app *App, code string) (*SessionOpenid, error) {
	session := &SessionOpenid{}
	m := map[string]string{
		"appid":      app.AppId,
		"secret":     app.Secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	err := doGetJson(app.client, app.Token, m, session, code2Session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
func GetAccessToken(app *App) (*AccessToken, error, ) {
	m := map[string]string{
		"appid":      app.AppId,
		"secret":     app.Secret,
		"grant_type": "client_credential",
	}
	res := &AccessToken{}
	err := doGetNotToken(app.client, m, res, getAccessToken)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//func GetUnlimitedWxaCodeFileName(data *QrcodeReq, accesstoken string) (interface{}, error) {
//	withByte, err := GetUnlimitedWxaCodeByte(, accesstoken)
//	randStr := string(RandString(10))
//	fileExtend := ".png"
//	pathRoot := "static/images/upload/"
//	filName := pathRoot + randStr + fileExtend
//	err = ioutil.WriteFile(filName, withByte, os.ModeAppend)
//	if err != nil {
//		fmt.Printf("err %s\n", err.Error())
//		return nil, err
//	}
//	return filName, nil
//}
//POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
func GetUnlimitedWxaCodeByte(app *App, data *QrcodeReq) ([]byte, error) {
	res, err := doGetBinary(app.client, app.Token, data, getUnlimited)
	return res, err
}

//POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
func SubMessageSend(app *App, data *SubMessageReq) (*SubMsgResp, error) {
	var subMsgResp = &SubMsgResp{}
	err := doPost(app.client, app.Token, data, subMsgResp, subscribeMessageSend)
	if err != nil {
		return nil, err
	}
	return subMsgResp, nil
}


func RandString(length int) []byte {
	base := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bs := []byte(base)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var res []byte
	for i := 0; i < length; i++ {
		n := r.Intn(len(bs))
		res = append(res, bs[n])
	}
	return res
}

func doGetBinary(client *resty.Client, token string, data interface{}, url string) ([]byte, error) {
	var m = make(map[string]string)
	res, err2 := json.Marshal(data)
	if err2 != nil {
		return nil, err2
	}
	err3 := json.Unmarshal(res, &m)
	if err3 != nil {
		return nil, err3
	}
	resp, err := client.R().SetQueryParams(map[string]string{
		"access_token": token,
	}).SetQueryParams(m).Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func doGetNotToken(client *resty.Client, data map[string]string, target interface{}, url string) error {
	resp, err := client.R().SetQueryParams(data).Get(url)
	if err != nil {
		return err
	}
	err = translateResponse(resp.Body(), target)
	if err != nil {
		return err
	}
	return nil
}
func doGetJson(client *resty.Client, token string, data interface{}, target interface{}, url string) error {
	var m = make(map[string]string)
	res, err2 := json.Marshal(data)
	if err2 != nil {
		return err2
	}
	err3 := json.Unmarshal(res, &m)
	if err3 != nil {
		return err3
	}
	resp, err := client.R().SetQueryParams(map[string]string{
		"access_token": token,
	}).SetQueryParams(m).Get(url)
	if err != nil {
		return err
	}
	err = translateResponse(resp.Body(), target)
	if err != nil {
		return err
	}
	return nil
}

func doPost(client *resty.Client, token string, data interface{}, target interface{}, url string) error {
	res, err := client.R().SetQueryParams(map[string]string{
		"access_token": token,
	}).SetBody(data).Post(url)
	if err != nil {
		return err
	}
	body := res.Body()
	fmt.Printf("do get body :\n%s\n", body)
	err = translateResponse(body, target)
	if err != nil {
		return err
	}
	return nil
}

func translateResponse(data []byte, target interface{}) error {
	var errInfo=&WxRepErr{}
	err := json.Unmarshal(data, errInfo)
	if err != nil {
		return err
	}
	if errInfo.Errcode != 0 {
		return errInfo
	}
	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	return nil
}
