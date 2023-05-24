package routers

import "go_vue_blog/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("settings/:name", app.ImageUploadView)
}
