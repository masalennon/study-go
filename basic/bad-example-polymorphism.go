package main

import (
	"fmt"
)

type worker struct {
	age, workingYear, baseSalary float64
	performance                  float64
	company                      string
}

func main() {
	taro := worker{
		age:         33,
		workingYear: 10,
		baseSalary:  50000,
		performance: 0.8,
		company:     "toyota",
	}
	hanako := worker{
		age:         28,
		workingYear: 5,
		baseSalary:  100000,
		performance: 1.90,
		company:     "google",
	}
	ichiro := worker{
		age:         40,
		workingYear: 15,
		baseSalary:  50000,
		performance: 1.30,
		company:     "sony",
	}
	fmt.Println(calculateIncome(taro.company))
}

func calculateIncome(w worker) int {
	switch w.company {
	case "toyota":
		return worker.baseSalary * (1.1 * worker.performance) * (worker.workingYear * 1.5)
	case "google":
		return worker.baseSalary * (8 * worker.performance)
	case "sony":
		return worker.baseSalary*(1.5*worker.performance) + (worker.baseSalary * worker.workingYear * 1.3)
	default:
		return fmt.Errorf("Error: %s", "dame-desu-yo")
	}
}
