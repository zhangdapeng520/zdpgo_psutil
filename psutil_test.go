package zdpgo_psutil

func getPsutil() *Psutil {
	return NewWithConfig(Config{
		Debug: true,
	})
}
