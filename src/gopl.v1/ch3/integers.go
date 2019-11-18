package main

import "fmt"

/*
	整数类型 分为 有符号和无符号类型

	有符号整数类型： int8，int16，int32和int64 四种大小不同
	无符号整数类型： uint8，uint16, uint32和uint64 四种大小不同  与上边上下对应

	与特定CPU 对应的整数类型 有符号整型 int 和 无符号整型 uint
	最广泛的应用是 int 对应的大小 32或64bit 不同的编译器即使在相同的硬件平台上也可能产生不同的大小

	Unicode字符类型 rune 和 int32 等价，通常表示一个 Unicode 码点。 两者可以互换使用
	byte 与 uint8 是等价类型。 但 byte 类型通常强调数值是原始数据而非一个小的整数

	uintptr 是无符号的整数类型，不指定具体 bit 大小，但足以容纳指针。 通常在底层编程时使用

	无论具体大小，int，uint和uintptr 是不同类型的兄弟类型，并且 int和int32 也是不同的类型(即使int的大小是32bit)，
	所以在使用时 需要显示的类型转换

	一个 n-bit 的有符号数的值域是从 $-2^{n-1}$到$2^{n-1}-1$。
	无符号整数的所有bit位都用于表示非负数，值域是0到$2^n-1$
	例如 int8 表示 -128到127 而uint8表示 0到255

	二元运算符的优先级顺序， 上到下递减  左优先
	*    /    %    <<    >>    &    &^
	+    -    |    ^
	==   !=   <    <=    >     >=
	&&
	||

	使用括号可以明确优先级顺序

	% 取模只用于整数间运算， 并且 在Go中，取模运算的符号与被取模数的符号总是一直的。所以 `-5%3` 与 `-5%-3` 的结果都是 -2
	/ 触发运算的截断 会向着 0 方向截断余数
*/
func main() {
	intOver()
	bitOp()
	reservedLoop()
}

func intOver() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) //255 0 1  255*255=0b1111111000000001 超过高位的bit位部分将被丢弃

	var i int8 = 127
	fmt.Println(i, i+1, i*i) //127 -128 1    127*127=0b11111100000001 超高高位的被截断
}

func bitOp() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// %08b 08表示至少8个字符宽度，不足前缀用0填充
	fmt.Printf("%08b\n", x) // 00100010
	fmt.Printf("%08b\n", y) // 00000110

	fmt.Printf("%08b\n", x&y)  // 00000010
	fmt.Printf("%08b\n", x|y)  // 00100110
	fmt.Printf("%08b\n", x^y)  // 00100100
	fmt.Printf("%08b\n", x&^y) // 00100000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // 1,5
		}
	}

	fmt.Printf("%08b\n", x<<1) // 01000100
	fmt.Printf("%08b\n", x>>1) // 00010001
}

func reservedLoop() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}


