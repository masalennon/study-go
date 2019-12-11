package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	//配列のスライス
	planets := [...]string {
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天皇星",
		"海王星",
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]

	//planetsの全要素を表すスライスを作成
	planets2 := planets[:]
	planets2[1] = "test"
	
	fmt.Println(terrestrial, gasGiants)
	fmt.Println(gasGiants[0])
	fmt.Println("planets2: ", planets2)

	fmt.Println("planets: ", planets)

	s := "abcde"
	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Println("b: ", b)
	}
	for _, r := range s {
		fmt.Println(r)
	}
	ss := "あ"
	for i := 0; i < len(ss); i++ {
    	fmt.Println("%v", ss[i]) // e3 81 82
	}
	fmt.Println('a')

	spacedPlanets := []string{"Mars ", "Earth ", " Venus"}
	hyperspace(spacedPlanets)
	fmt.Println(strings.Join(spacedPlanets, ""))

	type StringSlice []string
	
	//いきなりスライスを作る
	eightPlanets := []string {
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	sort.StringSlice(eightPlanets).Sort()

	fmt.Println(eightPlanets)


}

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}