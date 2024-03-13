package models

import "time"

type UserCollectModel struct {
	/*
		自定义第三张表，记录用户什么时候收藏了什么文章
	*/
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
