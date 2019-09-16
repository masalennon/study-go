package main

import (
	"fmt"
	"math"
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
		-28.0, 32.0, -31.0, -29.0, -29.0, 33.0,
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
	


}

