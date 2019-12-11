package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	fmt.Println(&v.X)
	fmt.Println(&v.Y)
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	var a rune = 'あ'
	fmt.Println(a)
	fmt.Println(&v.X)
    fmt.Println(&v.Y)
	
	v.Scale(10)
	fmt.Println(v.Abs())
}
