package global

import (
	"database/sql"
	"library/pkg/logger"
	"library/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db     *sql.DB
)
