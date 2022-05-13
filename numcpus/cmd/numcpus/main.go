package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_psutil/numcpus"
)

func main() {
	present, _ := numcpus.GetPresent()
	online, _ := numcpus.GetOnline()
	offline, _ := numcpus.GetOffline()
	possible, _ := numcpus.GetPossible()
	kernelMax, _ := numcpus.GetKernelMax()
	fmt.Printf("present:    %v\n", present)
	fmt.Printf("online:     %v\n", online)
	fmt.Printf("offline:    %v\n", offline)
	fmt.Printf("possible:   %v\n", possible)
	fmt.Printf("kernel_max: %v\n", kernelMax)
}
