package routers

import "time1043/gvb-server/api"

func (router RouterGroup) MenuRouter() {
	MenuApi := api.ApiGroupApp.MenuApi
	router.POST("menu", MenuApi.MenuCreateView)
}
