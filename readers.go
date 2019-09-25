package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") //13文字

	b := make([]byte, 8)
	for {
		fmt.Printf("b = %v\n", b)
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)		
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	test := "test"
	fmt.Println(test[0:3])
}
