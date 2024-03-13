package models

type MessageModel struct {
	MODEL

	SendUserID       uint      `json:"send_user_id" gorm:"primaryKey"` // 发送人id
	SendUserModel    UserModel `json:"-" gorm:"foreignKey:SendUserID"`
	SendUserNickName string    `json:"send_user_nick_name" gorm:"size:42"`
	SendUserAvatar   string    `json:"send_user_avatar"`

	RevUserID       uint      `json:"rev_user_id" gorm:"primaryKey"` // 接收人id
	RevUserModel    UserModel `json:"-" gorm:"foreignKey:RevUserID"`
	RevUserNickName string    `json:"rev_user_nick_name" gorm:"size:42"`
	RevUserAvatar   string    `json:"rev_user_avatar"`

	IsRead  bool   `json:"is_read" gorm:"default:false"` // 接收方是否查看
	Content string `json:"content"`                      // 消息内容
}
