package models

import (
	"github.com/astaxie/beego"
)

var message struct {
	toUserName   string `xml:ToUserName`
	fromUserName string
	createTime   int64
	msgType      string
	msgId        int64
}

var textMessage struct {
	message
	content string
}

func (tm textMessage) Println() {

}

const (
	text     = "text"
	image    = "image"
	voice    = "voice"
	video    = "video"
	location = "location"
	link     = "link"
	event    = "event"
)

type message interface {
	Print()
}
