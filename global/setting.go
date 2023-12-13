package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting    *setting.AppSettings
	MysqlSetting  *setting.MysqlSettings
	Logger        *logger.Logger
)
