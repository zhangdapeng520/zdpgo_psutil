# zdpgo_psutil

Go语言的系统工具库

项目地址：https://github.com/zhangdapeng520/zdpgo_psutil

## 版本历史

- v0.1.0 2022/4/8   基本功能
- v0.1.1 2022/4/22  新增：获取进程CPU信息
- v0.1.2 2022/4/22  BUG修复：修复一些依赖BUG
- v0.1.3 2022/5/13  新增：获取网卡信息
- v0.1.4 2022/5/13  新增：获取系统基本信息

## 使用示例
### 获取网卡信息
```go
package main

import "github.com/zhangdapeng520/zdpgo_psutil"

/*
@Time : 2022/5/13 15:34
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 获取网卡信息
*/

func main() {
	p := zdpgo_psutil.NewWithConfig(zdpgo_psutil.Config{Debug: true})
	p.GetNetworkInfo()
}
```

### 获取系统基本信息
```go
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
```