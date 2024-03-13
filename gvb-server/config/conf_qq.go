package config

import "fmt"

type QQ struct {
	AppId    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登录之后回调的地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppId == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppId, q.Redirect)
}
