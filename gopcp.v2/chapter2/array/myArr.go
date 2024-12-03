package main

import "fmt"

// array 存在长度初始化
var ipv4 = [...]uint8{192, 168, 0, 1}

// slice 没有长度
var ips = []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}

func main() {

	fmt.Println(ipv4)
	fmt.Println(ips)
	fmt.Println(ips[1])
	fmt.Println(ips[:cap(ips)])
	ips = append(ips, "192.168.0.4")
	fmt.Println(ips)

	newIps := make([]string, 100)
	fmt.Println(newIps)
	fmt.Println(cap(newIps))
	fmt.Println(len(newIps))

	intArr := make([]int8, 100)
	fmt.Println(intArr)
}
