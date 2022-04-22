# zdpgo_psutil

Go语言的系统工具库

项目地址：https://github.com/zhangdapeng520/zdpgo_psutil

## 版本历史

- 版本0.1.0 2022年4月8日 基本功能
- 版本0.1.1 2022年4月22日 新增：获取进程CPU信息
- 版本0.1.2 2022年4月22日 BUG修复：修复一些依赖BUG

## 使用示例
### 获取进程CPU信息
```go
package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_psutil"
)

func main() {
	p := zdpgo_psutil.New()
	info, err := p.GetProcessCpuInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println("进程1s内占用所有CPU比例：", info.AllCpuPercent)
	fmt.Println("进程1s内占用单个CPU比例：", info.SingleCpuPercent)
	fmt.Println("进程占用内存比例：", info.ProcessMemoryPercent)
	fmt.Println("创建的线程数：", info.ThredsNumber)
	fmt.Println("创建的goroutine数：", info.GoroutineNumber)
}
```
