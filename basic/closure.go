package main

import (
	"fmt"
)

func main() {
	test := closure()
	for i := 0; i < 10; i ++ {
		fmt.Println("closure", test(i))
		fmt.Println("normal ", normal(i))
	}

}

func closure() func(int) int {
	sum := 0 //closureの場合、ここの行が残り続けて、↓のfuncが←を参照し続けて行く感じかな
	return func(x int) int {
		sum += x
		return sum
	}
}

func normal(x int) int {
	amount := 0
	amount += x
	return amount;
}