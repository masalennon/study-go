package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	temperature := map[string]int {
		"Earth": 15,
		"Mars": -65,
		"Moon": 0,
	}
	temp := temperature["Earth"]
	fmt.Println("On averagem the earth is", temp)
	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Println("moon is :", moon)
	} else {
		fmt.Println("what is moon?")
	}

	//mapをコピーした時、同じものを参照していることを示す

	mapA := map[string]int {
		"a": 0,
		"b": 1,
		"c": 2,
	}
	mapB := mapA

	mapA["a"] = 1
	fmt.Println(mapA)
	fmt.Println(mapB)
	// countries := make(map[float64]int, 8)
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -29.0, 33.0, 33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
		fmt.Printf("frequency= %v\n", frequency[t])
		fmt.Printf("t= %v\n", t)
	}
	for t, num := range frequency {
		fmt.Printf("%+.2fの出現は%d回です。\n", t, num)
		fmt.Println(frequency)
	}

	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
	
	set :=  make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}
	if set[28.0] {
		fmt.Println("a memeber of set")
	}
	if set[-29.0] {
		fmt.Println("a memeber of set")
	}
	fmt.Println(set)

	text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	textLowerCase := strings.ToLower(text)
	trimedText := strings.Trim(textLowerCase, ",.?ー:")
	wordMap := strings.Fields(trimedText)
	// wordSet := make(map[string]bool) //mapを宣言する方法は複合リテラルかmakeの2通りのみ
	myCountWords := make(map[string]int)
	for _, t := range wordMap {
		myCountWords[t]++
	}
	for word, num := range myCountWords {
		if (num > 1) {
			fmt.Printf("%v: %v回 \n", word, num)
		}
	}

	//問題5の答え
	frequency1 := countWords(text)
	for word, count := range frequency1 {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}

}
func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `,."-`)
		frequency[word]++
	}
	return frequency
}

