package routers

import (
	"time1043/gvb-server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingApi

	// old: 127.0.0.1:8080/api/settings   127.0.0.1:8080/api/settings_email  2*5
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsUpdateView)

}
