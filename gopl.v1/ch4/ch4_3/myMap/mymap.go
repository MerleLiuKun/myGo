package main

import (
	"fmt"
	"sort"
)

/*
	哈希表 是一种巧妙并且实用的数据结构。 是一个 无序 的 key/value 对的集合，
	其中所有的 key 都是不同的，然后通过给定的 key 可以在常数时间复杂度内检索、更新和删除对应的 value。

	在 Go 中， 一个 map 就是一个哈希表的引用， map 类型可以写为 map[K]V，其中 K 和 V 分别对应 key 和 value。

	map 中所有的 key 都有相同的类型，所有的 value 也有着相同的类型，但是 key 和 value 之间可以是不同的类型。

	key 必须是支持 == 比较运算的数据类型。

	将浮点数类型作为 key 是一个不好的实现。

	使用内置的 make 函数可以创建一个 map。

	访问 map 的元素可以通过 key 直接访问。

	无法对 map 中的元素进行取址操作！
	因为 map 可能随着元素数量的增长而重新分配更大的内存空间，从而导致之前的地址无效。

	可使用 range 遍历 map，需要注意的是 map 的迭代顺序是不确定的。
	如果需要按顺序遍历 map 时，需要显式的对 key 进行排序。

*/

func main() {

	// make 内建函数
	ages := make(map[string]int)
	ages["alex"] = 21
	ages["merlin"] = 24
	ages["ikaros"] = 25
	fmt.Println(ages)

	// 使用字面值语法
	ages1 := map[string]int{
		"alex":   21,
		"merlin": 24,
		"ikaros": 25,
	}
	fmt.Println(ages1)

	// 创建空 map。
	blankMap := map[string]int{}
	fmt.Println(blankMap)

	// 访问 map 中的元素
	fmt.Println(getValue(ages, "merlin"))
	// 当 map 中没有对应的 key 时，也不会报错，会返回 value 值对应的 零值。
	fmt.Println(getValue(ages, "Ik"))

	// 删除 map 中的元素
	delKey(ages, "merlin")
	fmt.Println(ages)
	// 删除 不存在的元素
	delKey(ages, "Ik")
	fmt.Println(ages)

	simpleOp()
	lookupMap(ages)
	lookupMapBySort(ages)
	NoneMap()
	checkKey(ages)
}

func getValue(myMap map[string]int, key string) int {
	return myMap[key]
}

func delKey(myMap map[string]int, key string) {
	delete(myMap, key)
}

func simpleOp() {
	fmt.Println("Some operate for map value:")
	myMap := map[string]int{
		"ik": 1,
	}
	fmt.Println(myMap)

	myMap["ikaros"] = myMap["ikaros"] + 1
	fmt.Println(myMap)

	myMap["ik"] += 10
	fmt.Printf("Add 10 the ik is %v \n", myMap["ik"])
	myMap["ik"]++
	fmt.Printf("Call ++ the ik is %v \n", myMap["ik"])
}

func lookupMap(myMap map[string]int) {
	fmt.Println("Lookup map by rand sort:")
	for key, value := range myMap {
		fmt.Printf("%s\t%d\n", key, value)
	}
}

func lookupMapBySort(myMap map[string]int) {
	fmt.Println("Lookup map by sort:")
	//var names []string

	// 为了节省内存，可以创建一个容量与key的数量对应的 slice
	names := make([]string, 0, len(myMap))

	for name := range myMap {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, myMap[name])
	}
}

func NoneMap() {
	fmt.Println("Map is None: ")
	var ages map[string]int

	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)

	// 当 map 为零值时，无法直接赋值。
	// ages["age"] = 1
}

func checkKey(ages map[string]int) {
	// 检查 map 中的元素
	age, ok := ages["bob"]
	if !ok {
		fmt.Printf("Key bob not in map ages. %v\n", ok)
		fmt.Printf("Key bob zero is : %v\n", ages["bob"])
		ages["bob"] = 25
	}
	fmt.Printf("Key bob is : %v\n", ages["bob"])

	age, ok = ages["ikaros"]
	if !ok {
		fmt.Printf("Key ikaros not in map ages. %v\n", ok)
	}
	fmt.Printf("Key ikaros value is : %v\n", age)
}
