package main

import (
	"fmt"
)
//問題1
// 配列の各要素に引数で受け取った言葉をつけるプログラム
//問題2
// レシーバを持つメソッドを定義して、使用するプログラム
//問題3
// mapの要素の中で重複を調べて表示するプログラム











//配列の各要素に引数で受け取った言葉をつけるプログラム
func main() {
	animes := []string{"jojo", "naruto", "pokemon"}
	newAnimes := addWords("New ", animes...)
	fmt.Println(newAnimes)

	var jojo manga = "Jojo"
	newJojo := jojo.toAnime()
	fmt.Println(newJojo)
}

func addWords(prefix string, words ...string) []string {
	new := make([]string, len(words))
	for i := range words {
		new[i] = prefix + "" + words[i]
	}
	return new
}

type manga string
type anime string
func (m manga) toAnime() anime { //レシーバに処理を加えたいなら、引数ではなく、レシーバに型だけではなく（mangaだけではなく）インスタンスも（mも）定義することに注意
	animeNow := m + " is now anime"
	return anime(animeNow)
}