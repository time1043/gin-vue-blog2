package api

import (
	"time1043/gvb-server/api/images_api"
	"time1043/gvb-server/api/settings_api"
)

type ApiGroup struct {
	SettingApi settings_api.SettingApi
	ImagesApi  images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
