package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func (robot *simpleEN) Name() string {
	return robot.name
}

func (robot *simpleEN) Begin() (string, error) {
	return "Please input your name:", nil
}

func (robot *simpleEN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("Hello, %s! What can I do for you?", userName)
}

func (robot *simpleEN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "nothing", "bye":
		saying = "Bye!"
		end = true
		return
	default:
		saying = "Sorry, I didn't catch you."
		return
	}
}

func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("An error occurred: %s\n", err)
}

func (robot *simpleEN) End() error {
	return nil
}

func NewSimpleEN(name string, talk Talk) ChatBot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}
