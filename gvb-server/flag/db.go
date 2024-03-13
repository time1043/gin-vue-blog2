package flag

import (
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
)

func MakeMigrations() {
	//fmt.Println("xxx move ...") // go run main.go -db
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(&models.BannerModel{}, &models.TagModel{}, &models.MessageModel{}, &models.AdvertModel{},
			&models.UserModel{}, &models.CommentModel{}, &models.AdvertModel{}, &models.MenuModel{}, &models.MenuBannerModel{},
			&models.FadeBackModel{}, &models.LoginDataModel{})
	if err != nil {
		global.Log.Error("[ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ error ] 生成数据库表结构成功！")
}
