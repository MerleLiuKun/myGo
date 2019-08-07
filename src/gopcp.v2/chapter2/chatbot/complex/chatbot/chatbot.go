package chatbot

import "errors"

// 对话接口
type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

// 简易机器人的接口
type ChatBot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	ErrInvalidChatBotName = errors.New("invalid chatBot name")
	ErrInvalidChatBot     = errors.New("invalid chatBot")
	ErrExistingChatBot    = errors.New("existing chatBot")
)

var chatBotMap = map[string]ChatBot{}

func Register(bot ChatBot) error {
	if bot == nil {
		return ErrInvalidChatBot
	}
	name := bot.Name()
	if name == "" {
		return ErrInvalidChatBotName
	}
	if _, ok := chatBotMap[name]; ok {
		return ErrExistingChatBot
	}
	chatBotMap[name] = bot
	return nil
}

func Get(name string) ChatBot {
	return chatBotMap[name]
}
