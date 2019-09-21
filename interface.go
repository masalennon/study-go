package main

import (
	"fmt"
)
type SomeBehivor interface {
	DoSomething(arg string) string
}

// 構造体の定義
type StructA struct {
}

// インターフェースの実装
func (self *StructA) DoSomething(arg string) string {
	return arg + "a"
}

// ポリモフィズム
func main() {
	// インターフェース型の変数に格納できる
	var behivor SomeBehivor
	behivor = &StructA{}
	fmt.Println(behivor.DoSomething("A"))
}