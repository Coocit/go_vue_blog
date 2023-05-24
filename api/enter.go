package api

import (
	"go_vue_blog/api/images_api"
	"go_vue_blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
