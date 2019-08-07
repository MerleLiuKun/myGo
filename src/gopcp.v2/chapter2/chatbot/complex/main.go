package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopcp.v2/chapter2/chatbot/complex/chatbot"
	"os"
	"runtime/debug"
)

var chatBotName string

func init() {
	flag.StringVar(&chatBotName, "chatBot", "simple.en", "The chatBot's name for dialogue.")
}

func main() {
	flag.Parse()

	err := chatbot.Register(chatbot.NewSimpleEN("simple.en", nil))
	if err != nil {
		fmt.Println("Init en error")
	}
	err = chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	if err != nil {
		fmt.Println("Init cn error")
	}

	myChatBot := chatbot.Get(chatBotName)
	if myChatBot == nil {
		err := fmt.Errorf("Fatal error: Unsupport chatBot named %s\n", chatBotName)
		checkError(nil, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := myChatBot.Begin()
	checkError(myChatBot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(myChatBot, err, true)
	fmt.Println(myChatBot.Hello(input[:len(input)-1]))

	for {
		input, err = inputReader.ReadString('\n')
		if checkError(myChatBot, err, true) {
			continue
		}
		output, end, err := myChatBot.Talk(input)
		if checkError(myChatBot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
		if end {
			err = myChatBot.End()
			checkError(myChatBot, err, false)
			os.Exit(0)
		}
	}
}

func checkError(chatBot chatbot.ChatBot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatBot != nil {
		fmt.Println(chatBot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
