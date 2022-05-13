package zdpgo_psutil

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

type Psutil struct {
	Config *Config        // 配置对象
	Log    *zdpgo_log.Log // 日志对象
}

func New() *Psutil {
	return NewWithConfig(Config{})
}

func NewWithConfig(config Config) *Psutil {
	p := Psutil{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_psutil.log"
	}
	logConfig := zdpgo_log.Config{
		Debug:       config.Debug,
		OpenJsonLog: true,
		LogFilePath: config.LogFilePath,
	}
	if config.Debug {
		logConfig.IsShowConsole = true
	}
	p.Log = zdpgo_log.NewWithConfig(logConfig)

	// 配置
	p.Config = &config
	return &p
}
