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
	workers := []worker{
		ichiro, hanako, taro,
	}
	fmt.Println(calculateIncome(workers))
}

func calculateIncome(workers []worker) int {
	income := 0.0
	for _, worker := range workers {
		if worker.company == "toyota" {
			income += worker.baseSalary * (1.1 * worker.performance) * (worker.workingYear * 1.5)
		} else if worker.company == "google" {
			income += worker.baseSalary * (8 * worker.performance)
		} else if worker.company == "sony" {
			income += worker.baseSalary*(1.5*worker.performance) + (worker.baseSalary * worker.workingYear * 1.3)
		}
	}
	return int(income)
}
