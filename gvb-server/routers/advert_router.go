package routers

import "time1043/gvb-server/api"

func (router RouterGroup) AdvertRouter() {
	AdvertApi := api.ApiGroupApp.AdvertApi
	router.POST("advert", AdvertApi.AdvertCreateView)
	router.GET("advert", AdvertApi.AdvertListView)
	router.PUT("advert/:id", AdvertApi.AdvertUpdateView)
	router.DELETE("advert", AdvertApi.AdvertRemoveView)
}
