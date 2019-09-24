package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	a := make(map[string]bool)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}


	for i := 0; i < 5; i++ {
		gopherID := <-c //cに値が送信されてくるまで待機する
		fmt.Println("gopher", gopherID, "finished sleep.")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore...")
	c <- id
}