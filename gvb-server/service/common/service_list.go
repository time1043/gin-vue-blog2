package common

import (
	"gorm.io/gorm"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "create_at desc" // 默认按照时间往前排
	}

	query := DB.Where(model)
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model) // query会受上面查询id的影响 需要手动复位
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
