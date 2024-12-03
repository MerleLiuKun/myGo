package main

import "fmt"

//  一个简单的例子，将 列表 作为 Map 的键值

var m = make(map[string]int)

func main() {
	key1 := []string{"1", "2"}
	Add(key1)
	Add(key1)
	fmt.Printf("Now key %v count is %v", key1, Count(key1))

}

// 采用将列表转换为值的字符串作为 键。
func Key(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[Key(list)]++
}

func Count(list []string) int {
	return m[Key(list)]
}
