package main

import (
	"fmt"
)

func main() {

	// timmy := &person {
	// 	name: "Timothy",
	// 	age: 10,
	// }

	// timmy.superpower = "flying"
	// fmt.Printf("%+v\n", timmy)
	// superpowers := &[3]string{"fight", "invisibility", "super strength"}

	// fmt.Println(superpowers[0])
	// fmt.Println(superpowers[1:2])

	rebecca := person {
		name: "Rebecca",
		superpower: "imagination",
		age: 14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
	
	terry := &person {
		name: "Terry",
		age: 15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)

	magician := stats {
		level: 10,
		endurance: 11,
		health: 21,
	}


	player := character{name: "Matthias"}
	levelUp(&player.stats)
	fmt.Println(player)
	fmt.Println(magician)
	levelUp(&magician)
	fmt.Println(magician)

	var aaa float32
	fmt.Println(aaa)
	
}
type person struct {
	name, superpower string
	age int
}
func birthday(p *person) {
	p.age++
}
func (p *person) birthday() {
	p.age++
}

type stats struct {
	level int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats stats
}
