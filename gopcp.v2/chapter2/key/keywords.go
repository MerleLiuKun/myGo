package key

import "fmt"

type Name struct {
	firstName string
	lastName  string
}

type User struct {
	username string
	age      int
}

type MyWord struct {
	WName string
	user  User
	key   [2]string
}

func main() {
	a := Name{firstName: "ikaros", lastName: "kun"}
	fmt.Println(a.firstName)

	word := MyWord{
		WName: "MineCraft",
		user:  User{username: "Ikaros", age: 1},
		key:   [2]string{"player", "user"},
	}
	fmt.Println(word)

	fmt.Println("a\n a")
}
