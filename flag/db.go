package flag

import (
	"go_vue_blog/global"
	"go_vue_blog/models"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "MenuBanners", &models.MenuBannerModel{})
	//生成四张表的结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.BannerModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ success ] 生成数据库表结构成功!")
}
