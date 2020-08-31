package global

import (
	"github.com/damingerdai/blog-service/pkg/logger"
	"github.com/damingerdai/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
