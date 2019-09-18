package main

import (

)

func main() {
	type martian struct{}
	
	t = martian{}
	fmt.Println(t.talk())
}
func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

var t interface {
	talk() string
}

