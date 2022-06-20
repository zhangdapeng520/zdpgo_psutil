package zdpgo_psutil

import (
	"github.com/zhangdapeng520/zdpgo_log"
)

type Psutil struct {
	Config *Config        // 配置对象
	Log    *zdpgo_log.Log // 日志对象
}

func New(log *zdpgo_log.Log) *Psutil {
	return NewWithConfig(&Config{}, log)
}

func NewWithConfig(config *Config, log *zdpgo_log.Log) *Psutil {
	p := Psutil{}

	// 日志
	p.Log = log

	// 配置
	p.Config = config
	return &p
}
