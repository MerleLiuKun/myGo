package main

import "fmt"

// 事实上是 Map

var ipSwitch = map[string]bool{}

func main() {
	ipSwitch["192.168.0.1"] = true
	ipSwitch["192.168.0.2"] = false
	fmt.Println(len(ipSwitch))
	delete(ipSwitch, "192.168.0.2")
	fmt.Println(ipSwitch)
}
