package zdpgo_psutil

/*
@Time : 2022/5/13 14:40
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description: config 配置相关
*/

// Config 配置对象
type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`                 // 是否为debug模式
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"` // 日志路径
}