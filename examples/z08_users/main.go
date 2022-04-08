package main

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_psutil/libs/gopsutil/host"
)

func main() {
	// 获取终端用户信息
	users, _ := host.Users()
	for _, user := range users {
		data, _ := json.MarshalIndent(user, "", " ")
		fmt.Println(string(data))
	}
}
