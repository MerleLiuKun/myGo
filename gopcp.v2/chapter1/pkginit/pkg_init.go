package pkginit

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	fmt.Println(info)
	info = fmt.Sprintf("Os: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

var m = map[int]string{1: "A", 2: "B", 3: "C"}

var info string

func main() {
	fmt.Println(info)
}
