package models

import (
	"gorm.io/gorm"
	"os"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/ctype"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片hash值 用于判断图片是否重复
	Name      string          `json:"name" gorm:"size:38"`         // 图片名称
	ImageType ctype.ImageType `json:"image_type" gorm:"default:1"` // 图片存储类型 本地还是七牛云
}

// 钩子方法
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local { // 本地图片的删除 要删除本地存储
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
