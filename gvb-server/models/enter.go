package models

import "time"

type MODEL struct {
	ID       uint      `gorm:"primaryKey" json:"id"`                       // 主键ID
	CreateAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateAt time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP"` // 更新时间
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
