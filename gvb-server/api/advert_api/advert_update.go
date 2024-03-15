package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
)

func (AdvertApi) AdvertUpdateView(ctx *gin.Context) {
	id := ctx.Param("id")
	var cr AdvertRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	// 判断是否存在
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("该广告不存在", ctx)
		return
	}

	// 操作数据库添加
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", ctx)
		return
	}
	res.OkWithMessage("修改广告成功", ctx)
}
