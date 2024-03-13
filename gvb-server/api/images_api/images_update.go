package images_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请选择文件名称"`
}

func (ImagesApi) ImageUpdateView(ctx *gin.Context) {
	var cr ImageUpdateRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	// 操作数据库：先根据id查、再修改name
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", ctx)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("图片名称修改成功", ctx)
	return
}
