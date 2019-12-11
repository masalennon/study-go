package main

import (
	"fmt"
	"strings"
)

func main() {
	shout(martian{})
	shout(crator{})
}

type talker interface {
	talk() string
	chat() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	chatter := strings.ToLower(t.chat())
	fmt.Println(louder)
	fmt.Println(chatter)
}

type martian struct {}
type crator struct {}
func (m martian) talk() string {
	return "nack nack"
}
func (m martian) chat() string {
	return "wefoah aw;ehoi"
}
func (c crator) talk() string {
	return "crator"
}
func (c crator) chat() string {
	return "cratorrrr"
}