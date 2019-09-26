package main

import (
	"fmt"
)

func main() {
	a := sumOfCube(2, 5)
	fmt.Println(a)
}

func sumOfInt(n, m int) int {
    a := 0
    for ; n <= m; n++ {
        a += n
    }
    return a
}

func sumOfSquare(n, m int) int {
    a := 0
    for ; n <= m; n++ {
        a += n * n
    }
    return a
}

func sumOfCube(n, m int) int {
    a := 0
    for ; n <= m; n++ {
        a += n * n * n
    }
    return a
}