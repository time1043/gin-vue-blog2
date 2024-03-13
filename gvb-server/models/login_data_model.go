package models

type LoginDataModel struct {
	/*
		统计用户登录数据：id、用户id、用户昵称、用户token、登录设备、登录时间
	*/
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `json:"-" gorm:"foreignKey:UserID"`
	IP        string    `json:"ip" gorm:"size:20"` // 登录ip
	NickName  string    `json:"nick_name" gorm:"size:42"`
	Token     string    `json:"token" gorm:"size:256"`
	Device    string    `json:"device" gorm:"size:256"` // 登录设备
	Addr      string    `json:"addr" gorm:"size:64"`
}
