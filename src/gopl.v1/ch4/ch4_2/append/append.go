package main

import "fmt"

func main() {
	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d, len=%d\t%v\n", i, cap(y), len(y), y)
		x = y
	}

	var z []int
	z = appendIntMore(z, 1)
	fmt.Printf("cap=%d, len=%d\t%v\n", cap(z), len(z), z)
	z = appendIntMore(z, 2, 3)
	fmt.Printf("cap=%d, len=%d\t%v\n", cap(z), len(z), z)
	z = appendIntMore(z, z...)
	fmt.Printf("cap=%d, len=%d\t%v\n", cap(z), len(z), z)
}

func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1

	// 检查容量
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// 如果空间不足，创建一个新的 slice 用于保存新添加的元素
		zcap := zlen
		// 新 slice 的 容量为原容量的 2 倍。
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		// 复制源 slice 到新的 slice 中
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendIntMore(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// copy 新增的 slice 到 后方
	copy(z[len(x):], y)
	return z
}
