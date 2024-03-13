package settings_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/config"
	"time1043/gvb-server/core"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
)

func (SettingApi) SettingsUpdateView(ctx *gin.Context) {
	/*
		合并了五个接口
		PUT 127.0.0.1:8080/api/settings/site
		PUT 127.0.0.1:8080/api/settings/email
		PUT 127.0.0.1:8080/api/settings/qq
		PUT 127.0.0.1:8080/api/settings/qiniu
		PUT 127.0.0.1:8080/api/settings/jwt
	*/

	var cr SettingsUri
	err := ctx.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo // 获取用户输入json绑定
		err := ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.SiteInfo = info // 进行修改

	case "email":
		var info config.Email
		err := ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.Email = info

	case "qq":
		var info config.QQ
		err := ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.QQ = info

	case "qiniu":
		var info config.QiNiu
		err := ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.QiNiu = info

	case "jwt":
		var info config.Jwt
		err := ctx.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, ctx)
			return
		}
		global.Config.Jwt = info

	default:
		res.FailWithMessage("没有对应的配置项信息", ctx)
		return
	}

	// 最后修改yaml文件的配置
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	res.OkWith(ctx)

}
