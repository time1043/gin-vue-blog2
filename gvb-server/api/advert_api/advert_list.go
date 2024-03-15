package advert_api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
	"time1043/gvb-server/service/common"
)

func (AdvertApi) AdvertListView(ctx *gin.Context) {
	var cr models.PageInfo
	if err := ctx.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	// 判断referer：包含admin则全部返回
	referer := ctx.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(count, list, ctx)
}
