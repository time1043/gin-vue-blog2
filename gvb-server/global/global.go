package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time1043/gvb-server/config"
)

var (
	Config   *config.Config   // yaml配置信息 (指针 需要修改)
	Log      *logrus.Logger   // 日志配置
	DB       *gorm.DB         // gorm
	MysqlLog logger.Interface // mysql日志配置
)
