package main

import "flag"

var chatBotName string

func init()  {
	flag.StringVar(&chatBotName, "chatBot", "simple.en", "The chatBot's name for dialogue.")
}

func main() {

}
