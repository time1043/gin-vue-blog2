package images_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` // 图片路径
	Name string `json:"name"` // 图片名称
}

// ImageNameListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/image_name [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(ctx *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, ctx)
}
