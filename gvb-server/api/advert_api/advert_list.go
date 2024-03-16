package advert_api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
	"time1043/gvb-server/service/common"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo false "查询参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
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
