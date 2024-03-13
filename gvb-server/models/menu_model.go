package models

import "time1043/gvb-server/models/ctype"

type MenuModel struct {
	/*
		菜单表
	*/
	MODEL
	MenuTitle    string        `json:"menu_title" gorm:"size:32"`
	MenuTitleEn  string        `json:"menu_title_en" gorm:"size:32"`
	Slogan       string        `json:"slogan" gorm:"size:64"`                                                                     // slogan
	Abstract     ctype.Array   `json:"abstract" gorm:"type:string"`                                                               // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `json:"banners" gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID"` // 菜单的图片列表
	BannersTime  int           `json:"banners_time"`                                                                              // 菜单图片的切换时间 0表示不切换
	Sort         int           `json:"sort" gorm:"size:10"`                                                                       // 菜单的顺序
}
