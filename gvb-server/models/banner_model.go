package models

type BannerModel struct {
	MODEL
	Path string `json:"path"`                // 图片路径
	Hash string `json:"hash"`                // 图片hash值 用于判断图片是否重复
	Name string `json:"name" gorm:"size:38"` // 图片名称
}
