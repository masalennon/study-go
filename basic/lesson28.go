package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()

	time.Sleep(4 * time.Second)

	fmt.Println("sleepy")

}

func sleepyGopher() {
	time.Sleep(3 * time.Second)

	fmt.Println("...snore ...")
}