package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_vue_blog/global"
	"go_vue_blog/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	}

}
