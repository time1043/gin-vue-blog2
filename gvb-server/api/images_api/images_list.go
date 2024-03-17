package images_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
	"time1043/gvb-server/service/common"
)

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo false "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(ctx *gin.Context) {
	var cr models.PageInfo
	err := ctx.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	res.OkWithList(count, list, ctx)
	return
}
