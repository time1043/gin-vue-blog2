package models

import "time1043/gvb-server/models/ctype"

// UserModel 用户表
type UserModel struct {
	MODEL
	NickName       string           `json:"nick_name" gorm:"size:36"`                                                              // 昵称
	UserName       string           `json:"user_name" gorm:"size:36"`                                                              // 用户名 唯一
	Password       string           `json:"password" gorm:"size:128"`                                                              // 密码
	Avatar         string           `json:"avatar" gorm:"size:256"`                                                                // 头像
	Email          string           `json:"email" gorm:"size:128"`                                                                 // 邮箱
	Tel            string           `json:"tel" gorm:"18"`                                                                         // 手机号
	Addr           string           `json:"addr" gorm:"size:64"`                                                                   // 地址
	Token          string           `json:"token" gorm:"size:64"`                                                                  // 其他平台id
	IP             string           `json:"ip" gorm:"size:20"`                                                                     // ip地址
	Role           ctype.Role       `json:"role" gorm:"size:4;default:1"`                                                          // 权限 (1管理员 2普通登录者 3 游客 4封禁)
	SignStatus     ctype.SignStatus `json:"sign_status" gorm:"type=smallint(6)"`                                                   // 注册来源
	ArticleModels  []ArticleModel   `json:"-" gorm:"foreignKey:UserID"`                                                            // 发布的文章列表
	CollectsModels []ArticleModel   `json:"-" gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"` // 收藏的文章列表
}
