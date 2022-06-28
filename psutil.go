package zdpgo_psutil

type Psutil struct {
	Config *Config // 配置对象
}

func New() *Psutil {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Psutil {
	p := Psutil{}

	// 配置
	p.Config = config
	return &p
}
