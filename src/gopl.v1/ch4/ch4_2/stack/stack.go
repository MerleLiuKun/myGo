package main

import "fmt"

func main() {
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)
	fmt.Printf("stack: %v\n", stack)

	top := stack[len(stack)-1]
	fmt.Printf("Top is: %d\n", top)

	stack = remove(stack, 1)
	fmt.Printf("stack: %v\n", stack)

	stack = append(stack, 3, 4, 5, 6)
	fmt.Printf("stack: %v\n", stack)

	stack = removeNotKeepSort(stack, 2)
	fmt.Printf("stack: %v\n", stack)
}

// 删除 slice 中间的某个元素并保存原有的元素顺序
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 删除 slice 中间的某个元素 不保存原有顺序，将最后一位直接覆盖到被移除的元素上
func removeNotKeepSort(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1 ]
}
