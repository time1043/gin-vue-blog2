package models

type MenuBannerModel struct {
	/*
		自定义菜单和背景图片的连接表，方便排序
	*/
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"`
	Sort        int         `json:"sort" gorm:"size:10"`
}
