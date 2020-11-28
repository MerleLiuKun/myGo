package main

import (
	"fmt"
	"time"
)

/*
	结构体是一种聚合的数据类型，是由零个或者多个任意类型的值聚合而成的实体。
	每一项值都称为结构体的成员,

	结构体变量的成员可以通过 点操作符 访问，如 a.ID， 结构体是变量，他的成员也是变量。

	操作结构体的成员变量：
		1. 直接操作
		2. 对成员取地址，然后通过指针访问

	通常一行对应一个结构体的成员，名字在前，类型在后，如果相邻的成员类型相同可以合并到一行，

	成员的顺序有意义，成员也依赖于大写开头即可以导出的规则。

	命名为 S 的结构体不能有 S 类型的成员，但是 S 类型结构体可以包含 *S 指针类型的成员。

	如果结构体的全部成员是可以比较的，该结构体也可以使用 == 或者 ！= 比较。 将会比较所有成员。

*/

type Employee struct {
	ID int
	Name string
	Address string
	Dob time.Time
	Position string
	Salary int
	ManagerID int
}

type Point struct {
	X, Y int
}

var ik Employee

func main()  {
	op()
	structInitial()
	compare()
}

func op()  {
	ik = Employee{
		ID: 1,
		Name:"Ikaros",
		Address: "Beijing",
		Dob: time.Now(),

	}
	fmt.Println(ik.Address)

	ik.Salary -=5000
	fmt.Printf("Demoted salary, now is %d\n", ik.Salary)

	position := &ik.Position
	*position = "Senior" + *position
	fmt.Printf("New position, now is %s\n", ik.Position)

	var employeeOfTheMonth = &ik
	//employeeOfTheMonth.Position += " (proactive team player)" // 与 下方直接使用指针的方式一样
	(*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Printf("New position for ik, now is %s\n", ik.Position)
	fmt.Printf("New position for em, now is %s\n", employeeOfTheMonth.Position)

	fmt.Println(EmployeeByID(1).Position)

	// update by func
	id := ik.ID
	EmployeeByID(id).Salary = 0
	fmt.Printf("Fired, now salary is %d\n", ik.Salary)

	// 通过 结构体和 Map 来实现 集合 的数据结构 但应避免此用法，节约空间有限，而且语法复杂。
	/*
	seen := make(map[string]struct{})

	if _, ok := seen[s]; !ok{
		seen[s] = struct{}{}
	}
	*/
}

func structInitial()  {
	// 根据字面值语法指定
	// 顺序指定字面值初始化
	p:= Point{1,2}
	// 指定成员名称的初始化，推荐
	p1 := Point{X:3, Y:4}

	fmt.Println(p)
	fmt.Println(p1)

}

func EmployeeByID(id int) *Employee {
	/* ... */
	if id == ik.ID{
		return &ik
	}
	return nil
}

func compare()  {
	p:= Point{1, 2}
	q:= Point{2, 1}

	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	type address struct {
		hostname string
		port int
	}
	// 可比较的结构体可以用作 map 的 key
	hits := make(map[address]int)
	hits[address{"golang.org", 443}] ++
	fmt.Println(hits)
}

func nestedStruct()  {

	type Circle1 struct {
		X, Y, Radius int
	}

	type Wheel1 struct {
		X, Y, Radius, Spokes int
	}

	var w1 Wheel1
	w1.X = 8
	w1.Y = 8
	w1.Radius = 5
	w1.Spokes = 20

	// 抽离公共属性
	type Point struct {
		X, Y int
	}

	type Circle2 struct {
		Center Point
		Radius int
	}
	type Wheel2 struct {
		Circle Circle2
		Spokes int
	}

	var w2 Wheel2
	w2.Circle.Center.X = 8
	w2.Circle.Center.Y = 8
	w2.Circle.Radius = 5
	w2.Spokes = 20

	// 匿名成员
	type Circle3 struct {
		Point
		Radius int
	}
	type Wheel3 struct {
		Circle3
		Spokes int
	}

	var w3 Wheel3
	w3.X = 8
	w3.Y = 8
	w3.Radius = 5
	w3.Spokes = 20

	// 结构体字面值没有简短的表示法，但是并没有
	// w4 = Wheel3{8,8,5,20} Error
}