package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
	一个字符串是一个不可改变的字节序列。

	字符串可以包含任意的数据 包括 byte值 0， 但是通常用来包含人类可读的文本。

	文本字符串通常被解释为 采用 UTF8 编码的 Unicode 码点(rune)序列。

	内置的 len 函数可以返回一个字符串中的字节数目(注意，不是 rune 字符数目)，
	索引操作 s[i] 返回第 i 个字节的字节值，i 必须满足 0 <= i <= len(s)。

	访问越界将会导致 panic 异常。

	子字符串操作 s[i:j] 基于原始 s 字符串的 第 i 个字节开始到 第 j 个字节(不包括 j 本身)生成一个新的字符串。 左开右闭。
	生成的字符串包含 j-i 个字节。
	这里的 i 和 j 均可省略。忽略时采用 0 作为开始，采用 len(s) 作为结束的位置。

	+ 操作符可以将两个字符串拼接到一起构造一个新的字符串。

	字符串可以用 == 和 < 等进行比较: 通过逐个字节比较完成。

	字符串的值是不可变的： 一个字符串包含的字节序列永远不会被改变。但是可以给一个字符串变量分配一个新的字符串值。比如追加。
	所以修改字符串内部数据的操作是禁止的。 s[0] = 'L' 会编译错误。

	标准库中有四个包对字符串处理尤为重要： bytes、strings、stconv 和 unicode 包。
	strings 包提供了许多如字符串的查询，替换，比较，截断，拆分和合并等功能的函数
	bytes 包也提供了类似功能的函数，是针对与字符串有着相同结构的 []byte 类型，
	因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下 使用 bytes.Buffer 类型会更加有效。

	strconv 包提供了布尔型、整数型、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。

	unicode 包提供了 IsDigit、IsLetter、IsUpper 和 IsLower 等功能函数，用于给字符分类。

	字符串和数值之间的转换也比较常见。 strconv 提供了这类转换功能。

*/

func main() {
	s := "Hello Ikaros"
	fmt.Println(len(s))                // 12
	fmt.Println(s[0], s[7])            // 72 107
	fmt.Printf("%q, %q\n", s[0], s[7]) // 'H', 'K'

	//c := s[len(s)]
	//fmt.Println(c)

	s1 := "中国"
	fmt.Println(len(s1))      // 6
	fmt.Printf("%q \n", s[1]) // 'e'

	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(s[:])

	fmt.Println("Good" + s[5:])

	s2 := "left foot"
	fmt.Println(s2)
	fmt.Println(&s2)
	t := s2
	s2 += ", right foot"
	fmt.Println(s2)
	fmt.Println(&s2) // 地址没有变化 但是该地址赋予了新的字符串。
	fmt.Println(t)
	fmt.Println(&t)

	fmt.Println(HasPrefix(s, "Hello"))
	fmt.Println(HasSuffix(s, "ros"))
	fmt.Println(Contains(s, "Ika"))

	unicodeHandler()
	unicodeToRune()
	unicodeToString()

	numberToString()
	stringToNumber()
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func unicodeHandler() {
	s := "Hello, 刘坤"

	fmt.Println(len(s))                    // 13 13个字节
	fmt.Println(utf8.RuneCountInString(s)) // 9  9个 unicode 字符

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:]) // 函数会获得一个 unicode 字符
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// range 循环处理字符串的时候 会自动隐式解码 UTF8 字符串。
	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	n := 0
	for _, _ = range s { // _ 可以省略
		n++
	}
	fmt.Println(n)

}

func unicodeToRune() {
	s := "刘坤"
	fmt.Printf("% x\n", s) // % x 会在每个十六进制数前插入一个空格

	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println(string(r)) // 进行 string 转换时，会进行 UTF8 编码
}

func unicodeToString() {
	// 将一个整数转型为字符串的意思是生成 只包含对应 Unicode 码点字符的 UTF8 字符串。
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))

	// 如果对应码点的字符是无效的 则用 \uFFFD 无效字符作为替换
	fmt.Println(string(123456))
}

func numberToString() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// 进制转化
	fmt.Println(strconv.FormatInt(int64(x), 2))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
}

func stringToNumber() {
	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(x)

	y, err := strconv.ParseInt("123", 10, 64)  // 64表示整数大小，int64。 0-> int。
	fmt.Println(y)
}
