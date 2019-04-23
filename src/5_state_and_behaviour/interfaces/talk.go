package interfaces

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", 3)
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func TestInterfaces() {
	shout(martian{})
	shout(laser(2))
}
