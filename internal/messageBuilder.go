package internal

import (
	"github.com/fatih/color"
)

type Message struct {
	message string
}

func NewMessage(col color.Attribute, message string) *Message {
	return &Message{
		message: color.CyanString("[GD] ") + color.New(col).Sprint(message),
	}
}

func (msg *Message) ThenColor(col color.Attribute, message string) *Message {
	msg.message += " " + color.New(col).Sprint(message)
	return msg
}

func (msg *Message) ThenColorStyle(col color.Attribute, style color.Attribute, message string) *Message {
	msg.message += " " + color.New(col).Add(style).Sprint(message)
	return msg
}

func (msg *Message) String() string {
	return msg.message
}