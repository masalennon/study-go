package main

import (
	"strings"
	"fmt"
)
type talker interface {
	talk() string
}
//talkを満たす型ならなんでも入れることができる
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func main() {
	shout(martian{})
	shout(&martian{})
}