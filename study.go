package main

import (
	"fmt"
	"math/rand"
	"strings"
	"math/big"
)

func main() {
	// 変数の宣言
	var weight = 61
	fmt.Println("Hello, Go")
	fmt.Println(weight)
	weight -= 2
	fmt.Println(weight)
	weight *= 2
	fmt.Println(weight)

	// varを省略できる
	height := 180
	fmt.Println(height)

	//ランダムに変数を生成（0-9）
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	var num2 = rand.Intn(401000001-56000000) + 56000000
	fmt.Println(num2)

	//Lesson 2 練習問題
	const distance = 56000000
	const days = 28
	const hoursPerDay = 24
	fmt.Println("distance = ", distance / (days * hoursPerDay), "km/h")
	fmt.Printf("1: %v 2: %v", "aew", "afew")

	fmt.Println("you are in a dark cave")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("exit the cave", exit)
	

	//forループ、goにwhileループはない
	var count = 5
	for count > 0 {
	  fmt.Println(count)
	  count--
	}

	// switch文 varを省略した書き方をすることで、変数を宣言できる
	switch num3 := rand.Intn(3); num3 {
	case 0:
	  fmt.Println("Space Adventures")
	case 1:
	  fmt.Println("SpaceX")
	default:
	  fmt.Println("random")
	}

	//Lesson 5 火星行きのチケット
	fmt.Println("===============================Lesson 5　火星行きのチケット======================================")
	
	company := ""
	travelType := ""
	secondsPerDay := 60*60*24
	
	for count := 10; count > 0; count-- {
		speed := rand.Intn(15) + 16
		price := speed + 20
        days := 62100000 / speed / secondsPerDay
		if tripType := rand.Intn(2); tripType == 0 {
		  travelType ="Round-trip"
		  price *= 2
		} else {
		  travelType ="One-way"
		}

	  switch rand.Intn(4) {
		case 0:
			company = "SpaceX"
	    case 1:
			company = "Space Adventures"
	    case 2:
			company = "Virginia"
	  }
	  fmt.Printf("%v %v %v $ %v \n", company, days, travelType, price)

	}

	// %fをつけて出力の整形 5は幅、2は精度を表す。幅（小数点を含めた文字数）が5になるように、スペースがつく。％の左に0がある時は、足りない場合は0で埋める
	fmt.Printf("%05.2f \n", 0.3333333)	
	
	// Lesson 6 実数
	fmt.Println("===============================Lesson 6 実数======================================")
	saving := 0.0;
	

	for  saving < 1.00 {
		switch rand.Intn(3) {
		case 0:
			saving += 0.05
		case 1:
			saving += 0.10
		case 2:
			saving += 0.25
		}
		fmt.Printf("current saving = %5.2f \n", saving)
	}
	

	//Lesson 7 整数

	fmt.Println("===============================Lesson 7 整数======================================")

	var year int = 2018

	var month uint = 2

	
	fmt.Printf("%T型： %v", year, month) //%Tで型を調べる
	fmt.Printf("%T型： %[1]v \n", month) //第一引数を意味する[1]を指定

	//ビットを表示する
	var green uint8 = 3
	fmt.Printf("%08b\n", green)

	fmt.Println("===============================Lesson 7 練習問題======================================")
	
	newSaving := 0 //これを0にすると型がfloat64になり、%が使えなくなる

	for newSaving < 200 {
		switch rand.Intn(3) {
		case 0:
			newSaving += 5
		case 1:
			newSaving += 10
		case 2:
			newSaving += 25
		}

		dollars := newSaving / 100
		cents := newSaving % 100
		fmt.Printf("$ %v %v dollars %v cents\n", newSaving, dollars, cents)
	}


	var earth uint64 = 56e6
	fmt.Println(earth)

	lightSpeed := big.NewInt(299792)
	distance1 := new(big.Int)
	distance1.SetString("24000000000000000000000000", 10)

	seconds := new(big.Int)
	seconds.Div(distance1, lightSpeed)

	






	
	
	
	
}
