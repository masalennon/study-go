package main

import (
	"fmt"
)

func main() {

	studentA := person{
		name: "Taro",
		age:  20,
	}

	fmt.Println(studentA)
	studentA.selfIntro()

	taro := japanese{
		language: "Japanese",
		country:  "Japan",
		person:   studentA,
	}
	ichiro := newJapanese("Ichiro", 33)
	fmt.Println(ichiro)

	fmt.Println(taro)
	fmt.Println(taro.person.name)

	nara := wheatherReport{
		temperature: temperature{high: 21.2, low: 11.0},
		humidity: humidity{high: 55, low: 33},
	}
	fmt.Println(nara.humidity.high)
	fmt.Println(nara.temperature.high)

	tem := newTemperature(21, 2)
	fmt.Println(tem)
}

func newJapanese(name string, age int) person {
	return person{name, age}
}

func newTemperature(high, low float64) temperature {
	return temperature{high, low}
}

type person struct {
	name string
	age  int
}

type japanese struct {
	language string
	country  string
	person   person
}

func (p person) selfIntro() {
	fmt.Printf("My name is %v. I'm %v years old.\n", p.name, p.age)
}

type temperature struct {
	high, low float64
}

type humidity struct {
	high, low float64
}

type location struct {
	latitude, longitude float64
}

type wheatherReport struct {
	temperature temperature
	humidity humidity
	location location
}

