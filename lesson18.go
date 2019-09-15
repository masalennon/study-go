package main

import (
	"fmt"
)

func main() {
	books := []string{"Harry", "Dragonball", "NARUTO"}
	books = append(books, "BLEACH")
	fmt.Println(books)
	dump("books", books)
	books1 := append(books, "SLAMDUNK")
	dump("books1: ", books1)

	//容量6の配列に7、8個目を追加しようとしたので大きさが2倍の配列がコピーされた。元の配列とは違うものになっている
	books2 := append(books1, "HUNTER", "courage to be disliked")
	dump("books2: ", books2)
	books2[1] = "---"
	dump("books2: ", books2)
	dump("books1: ", books1)
	books3 := books1[0:4:4] //[0:4]とすると。SLAMDUNKが上書きされてしまう
	dump("books3: ", books3)
	books4 := append(books3, "JOJO")
	dump("books4: ", books4)
	fmt.Println(books1)

	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := transform("New", planets...)
	fmt.Println(newPlanets)

}

func dump(lable string, slice []string) {
	fmt.Println(lable, len(slice), cap(slice), slice)
}

//可変個引数を宣言するには、省略記号（...）を最後のパラメータに前置して使う
func transform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + "" + worlds[i]
	}
	return newWorlds
}