package main

import (
	"fmt"
)

func main() {
	sample := 42
	fmt.Println(&sample)
	fmt.Println(*&sample)
	sampleAddress := &sample
	fmt.Println(*sampleAddress)

	canada := "Canada"
	var home *string
	fmt.Printf("homeは %Tです。\n", home)
	home = &canada
	fmt.Println(*home)

	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println("*administrator: " + *administrator)

	*administrator = "Maj. Gen. Charles Frank Bolden Jr." //boldenの値をポインタを使って間接的に変更
	fmt.Println("bolden: " + bolden)

	major := administrator
	fmt.Println(*major)
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(*major)
	fmt.Println(administrator == major) //同じメモリアドレス

	charles := *major //別のメモリアドレスになる
	fmt.Println("charles: " + charles)
	fmt.Println(&charles, &*major)
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)
	
	
}