package main

import "fmt"

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type ChatBot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

type myTalk string // 此处的 string 啥作用？

func (talk *myTalk) Begin() (string, error) {
	return "Begin", nil
}

func (talk *myTalk) ReportError(err error) string {
	panic("implement me")
}

func (talk *myTalk) End() error {
	panic("implement me")
}

func (talk *myTalk) Hello(userName string) string {
	fmt.Println("Begin to hello....")
	return "Ikaros"
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	fmt.Println("I heard ", heard)
	return "Hei, My merlin", false, nil
}

func (talk *myTalk) Name() string {
	return "Ikaros"
}

// 组合结构体
type simpleCN struct {
	name string
	talk Talk
}

func main() {
	var talk ChatBot = new(myTalk)

	name := talk.Hello("Ik")
	st, end, err := talk.Talk(name)
	if err != nil {
		fmt.Println(end)
	}
	fmt.Println(st)

	sa, err := talk.Begin()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sa)

	a := simpleCN{"Ikaros", new(myTalk)}
	fmt.Println(a.talk.Hello("Me"))
}
