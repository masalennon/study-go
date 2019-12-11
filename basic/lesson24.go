package main

import (
	"fmt"
)
//簡単なインターフェースを定義して使用するプログラム

func main() {
	b = dog{}
	fmt.Println(b.bark())
	b = cat(3)
	c = catRobot(90)
	fmt.Print(c.number())
}


type dog struct {}
type cat string

var b interface {
	bark() string
}
// ①インターフェイス型を宣言する
// ②その型を引数に取るメソッドを宣言する
// ③ある型を宣言する
// ④その型をレシーバに持つメソッドを定義する
// ⑤そのメソッドの引数に、宣言したある型を使う
func (d dog) bark() string {
	return "woof woof"
}

func (c cat) bark() string {
	return "meow" + string(c)
}


var c interface {
	number() int
}

type catRobot int

func (cr catRobot) number() int {
	return 123 + int(cr)
}

// func main() {
// 	t = martian{}
// 	fmt.Println(t.talk())
// 	t = laser(3)
// 	fmt.Println(t.talk())
// 	t = dog{}
// 	fmt.Println(t.talk())
// }     

// func (d dog) talk() string {
// 	return "....."
// }
// type dog struct {}
// type laser int

// func (l laser) talk() string {
// 	return "pew"
// }
// //
// var t interface {
// 	talk() string
// }
