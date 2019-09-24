package main

import (
	"time"
	"fmt"
)

func main() {
	// go sleepyGopher()

	for i := 0; i < 5; i++ {
		sleepyGopherInt(i)
	}

	time.Sleep(4 * time.Second)
	fmt.Println("... snore 2...")
	
}

// func sleepyGopher() {
// 	time.Sleep(4 * time.Second)

// 	fmt.Println("... snore ...")
// }
func sleepyGopherInt(id int) {
	time.Sleep(4 * time.Second)

	fmt.Println("... snore ...", id)
}