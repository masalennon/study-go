package main

import (
	"fmt"
)


//配列の各要素に引数で受け取った言葉をつけるプログラム
func main() {
	animes := []string{"Naruto", "Nichijou", "Jojo", "Dragon ball"}
	fmt.Println(animes)
	for _, word := range animes {
		fmt.Println("New " + word )
	}
}
