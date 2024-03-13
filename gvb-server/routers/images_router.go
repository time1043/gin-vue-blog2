package routers

import "time1043/gvb-server/api"

func (router RouterGroup) ImagesRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.GET("images", ImagesApi.ImageListView)      // 127.0.0.1:8080/api/images?limit=5&page=2
	router.POST("images", ImagesApi.ImageUploadView)   // 127.0.0.1:8080/api/images  图片
	router.DELETE("images", ImagesApi.ImageRemoveView) // (Body row json) "id_list": [1,2,3,4]

}
