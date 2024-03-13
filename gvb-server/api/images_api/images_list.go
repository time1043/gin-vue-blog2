package images_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
	"time1043/gvb-server/service/common"
)

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
