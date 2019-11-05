package main

import "fmt"

type Website struct {
	name   string
	length int
	url    string
}

var site = Website{
	name:   "StudyGolang",
	length: 1024,
	url:    "https://studygolang.com/",
}

func main() {
	// 通用
	fmt.Printf("%v\n", site)  // 相应值的默认格式
	fmt.Printf("%+v\n", site) // 输出加上字段名
	fmt.Printf("%#v\n", site) // 值的Go语言表示
	fmt.Printf("%T\n", site)  // 类型的Go语言表示
	fmt.Printf("%%\n")        // 字面上的 %
	// 布尔型
	fmt.Printf("%t\n", true) // true 或者 false 的单词
	// 整型
	fmt.Printf("%b\n", 1024)     // 二进制表示
	fmt.Printf("%c\n", 11111111) // Unicode码点对应的字符
	fmt.Printf("%d\n", 10)       // 十进制表示
	fmt.Printf("%o\n", 8)        // 八进制表示
	fmt.Printf("%#o\n", 8)       // 带 0 前缀的八进制表示
	fmt.Printf("%q\n", 22)       // 由 Go 语言转义的 单引号围绕的字符字面值
	fmt.Printf("%x\n", 1223)     // 十六进制表示，字母形式为小写
	fmt.Printf("%X\n", 1223)     // 十六进制表示，字母形式为大写
	fmt.Printf("%U\n", 1223)     // Unicode格式 U+xxx
	// 浮点数和复数
	fmt.Printf("%b\n", 12.34)    // 无小数部分的，指数为二的幂的科学计数法
	fmt.Printf("%e\n", 12.345)   // 科学计数法 小写e
	fmt.Printf("%E\n", 12.34567) // 科学计数法 大写E
	fmt.Printf("%f\n", 12.3456)  // 有小数点而无指数
	fmt.Printf("%F\n", 12.3456)  // 与 f 一致
	fmt.Printf("%g\n", 12.3456)  // 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%G\n", 12.3456)  // 根据情况选择 %E 或 %F 以产生更紧凑的（无末尾的0）输出
	// x, X 的作用？

	// 字符串
	fmt.Printf("%s\n", "qwer")   // 输出字符串表示（string类型或[]byte)
	fmt.Printf("%q\n", "qwer")   // 双引号围绕的字符串
	fmt.Printf("%x\n", "asdzxc") // 十六进制，小写字母，每字节两个字符
	fmt.Printf("%X\n", "asdzxc") // 十六进制，大写字母，每字节两个字符
	v := "123"
	fmt.Printf("%p\n", &v) // 十六进制表示，前缀 0x

	sli := []string{"1", "2"}
	// 切片
	fmt.Printf("%p\n", sli) // 十六进制表示首位元素的地址，前缀 0x
}
