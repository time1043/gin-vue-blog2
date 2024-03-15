package routers

import "time1043/gvb-server/api"

func (router RouterGroup) AdvertRouter() {
	AdvertApi := api.ApiGroupApp.AdvertApi
	router.POST("advert", AdvertApi.AdvertCreateView)
	//router.GET("advert", AdvertApi.ImageUploadView)
	//router.DELETE("advert", AdvertApi.ImageRemoveView)
	//router.PUT("advert", AdvertApi.ImageUpdateView)
}
