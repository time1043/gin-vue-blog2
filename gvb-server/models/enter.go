package models

import "time"

// 数据库表的基础字段
type MODEL struct {
	ID       uint      `gorm:"primaryKey" json:"id"`                          // 主键ID
	CreateAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP(3)"` // 创建时间
	UpdateAt time.Time `json:"update_at" gorm:"default:CURRENT_TIMESTAMP(3)"` // 更新时间
}

// 分页展示
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

// 删除时发的请求参数
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
