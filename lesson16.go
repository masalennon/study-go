package main

import (
	"fmt"
)


func main() {

	var planets [8]string

	planets[0] = "水星"
	planets[1] = "金星"
	planets[2] = "地球"

	earth := planets[2]

	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	// panic runtime errorになる
	// i := 8
	// planets[i] = "冥王星"
	// pluto := planets[i]
	// fmt.Println(pluto)


	//複合リテラル構文
	dwarfs := [5]string{"ケレス", "冥王星", "ハウメア", "マケマケ", "エリス"}

	//宣言する配列の要素の数をコンパイルに数えさせるなら...を使う
	fruits := [...]string {
		"apple",
		"banana",
		"orange",
		"grape",
		"peach", //最後にもカンマが必要
	}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
	fmt.Println(dwarfs)
	fmt.Println(len(fruits))

	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"

	// board[1]にある列の数の分だけループ
	for column := range board[1] {
		board[1][column] = "p"
		fmt.Println(column)
		fmt.Println("board[1]:", board[1])
		fmt.Println(board[column])
	}
	fmt.Print(board)

}
