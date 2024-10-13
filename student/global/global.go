package global

import (
	"gorm.io/gorm"
	"student/pkg/logger"
	"student/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db     *gorm.DB
)
