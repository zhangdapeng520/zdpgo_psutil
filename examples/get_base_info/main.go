package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil"
)

/*
@Time : 2022/5/13 15:34
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 获取网卡信息
*/

func main() {
	p := zdpgo_psutil.NewWithConfig(zdpgo_psutil.Config{Debug: true})
	info, err := p.GetBaseInfo()
	if err != nil {
		return
	}
	fmt.Println(info)
}
