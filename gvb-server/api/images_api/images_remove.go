package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
)

func (ImagesApi) ImageRemoveView(ctx *gin.Context) {
	// 删除需要传入请求参数
	var cr models.RemoveRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	// 文件不存在
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", ctx)
		return
	}
	// 删除数据库数据 批量删除
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 张图片", count), ctx)
}
