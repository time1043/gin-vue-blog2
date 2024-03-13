package settings_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingApi) SettingsInfoView(ctx *gin.Context) {
	/*
		合并了五个接口
		GET 127.0.0.1:8080/api/settings/site
		GET 127.0.0.1:8080/api/settings/email
		GET 127.0.0.1:8080/api/settings/qq
		GET 127.0.0.1:8080/api/settings/qiniu
		GET 127.0.0.1:8080/api/settings/jwt
	*/
	
	var cr SettingsUri
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, ctx) // 127.0.0.1:8080/api/settings/site
	case "email":
		res.OkWithData(global.Config.Email, ctx) // 127.0.0.1:8080/api/settings/email
	case "qq":
		res.OkWithData(global.Config.QQ, ctx)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, ctx)
	case "jwt":
		res.OkWithData(global.Config.Jwt, ctx)
	default:
		res.FailWithMessage("没有对应的配置项信息", ctx)
	}

}
