package api

import (
	"time1043/gvb-server/api/advert_api"
	"time1043/gvb-server/api/images_api"
	"time1043/gvb-server/api/menu_api"
	"time1043/gvb-server/api/settings_api"
)

type ApiGroup struct {
	SettingApi settings_api.SettingApi
	ImagesApi  images_api.ImagesApi
	AdvertApi  advert_api.AdvertApi
	MenuApi    menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
