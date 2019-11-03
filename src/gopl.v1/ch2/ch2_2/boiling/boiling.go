package boiling

import "fmt"

const boilingF = 212.0 // 包级别常量 可以在包对应的源文件中使用

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9 // f,c 函数级别变量 只能在函数内部使用
	fmt.Printf("Boiling point = %g °F or %g°C\n", f, c)
}
