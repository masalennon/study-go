package main

import (
	"fmt"
	"math/rand"
	"strings"
	"math/big"
	"math"
	"strconv"
	"time"
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
		case 0, 4, 5: //switchの候補を並べて書ける
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

	const distanceSpace = 236000000000000000
	const lightSpeedPerSocond = 299792
	const daysPerYear = 365
	const lightSpeedInt = 299792

	const yearsTillSpace = distanceSpace / lightSpeedInt / daysPerYear / lightSpeedPerSocond

	fmt.Printf("years until the space: %v", yearsTillSpace)

	grade := 'A'
	fmt.Printf("grade: %v", grade)
	peace := "shalom"
	peace = "sama"

	fmt.Printf("peace: %v", peace)

	//1文字ずつ表示

	message := "shalom"
	for i := 0; i < 6; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}

	c := 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c\n", c)

	alienMessage := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(alienMessage); i++ {
		a := alienMessage[i]
		if a >= 'a' && a <= 'z' { //アルファベットのみが置換の対象
			a = a + 13 //13ずらす
			if a > 'z' { //zを超えてしまったら26引いて戻す
				a = a -26
			}
		}






		fmt.Printf("%c", a)
	}
	// moji := "aiueo"
	// moji[0] = 'e' Stringはimmutable
	// fmt.Printf("%v ", moji)
	
	countdown := "only " + "10 seconds"
	fmt.Printf("\n%v", countdown)
	
	v := 42
	if v > 0 && v < math.MaxUint8 {
		vuint := uint8(v)
		fmt.Print(vuint)
	}

	count1 := 10
	str := "\nlast " + strconv.Itoa(count1) + " seconds \n"
	fmt.Printf(str)

	// countdown1 := 10
	// countdown1 = 0.5 エラー　

	// countdown1 = fmt.Sprintf("%v ", countdown1)　エラー


	answer := true
	var oneZero int
	if answer {
		oneZero = 1
	} else {
		oneZero = 0
	}
	fmt.Print(oneZero)
	fmt.Println("")

	//==================================配列の中身を一つ一つ表示=====================================

	words := []string{"apple", "orange", "pineapple"}
	for _, word := range words {
		fmt.Printf("%v ", word)
	}


	// =================================定義したメソッドを使用=====================================
	var kelvin0978 kelvin = 294.0
	celsius2 := kelvinToCelsius(kelvin0978)
	fmt.Print(kelvin0978, "K is ", celsius2, "C.")

	ff := celsiusToFahrenheit(kelvinToCelsius(kelvin(0)))
	fmt.Print("k = ", ff)

	//===================================新しい型を定義============================================
	type celsius1 float64
	const degrees = 20
	var temperature celsius1 = degrees
	temperature += 10

	var warmUp float64 = 10
	// temperature += warmUp 型の不一致でエラーになる

	temperature += celsius1(warmUp) //型を変換する必要がある

	var cel celsius1 = 20
	var fah farenheit = 20
	// if cel == fah { // invalid operation　タイプの不一致　不正な演算

	// cel += fah // invalid oepration
	fmt.Print("\n", cel, fah)



	var kel kelvin = 294.0
	var kToC celsius

	kToC = kelvinToCelsius(kel)

	kToC = kel.celsius()

	fmt.Printf("\nkToC: %v\n", kToC)


	//=====================ファーストクラス関数==========================

	sensor := fakeSensor
	fmt.Println(sensor())
	
	sensor = realSensor
	fmt.Println(sensor())

	// ==========================無名関数==============================

	f := func(message string) {
		fmt.Println(message)
	}
	f("this is an anonymous method.")

	func() {
		fmt.Println("Functions anonymous")
	}()


}

type farenheit float64
type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	c := k - 273.15
	return celsius(c)
}

//カスタムの型を使う変数
func celsiusToFahrenheit(c celsius) farenheit {
	f := farenheit(c * 9.0 / 5.0) + 32.0 //型変換が必要
	return f
}

//レシーバを持ったメソッド

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (k kelvin) celsius() celsius {
	 return celsius(k - 273.15)
}
func (f farenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (c celsius) farenheit() farenheit {
	return farenheit((c * 9.0 / 5.0) + 32.0)
} 
func (f farenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}
func (k kelvin) farenheit() farenheit {
	return k.celsius().farenheit()
}


func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}

type getRowFn func(rows int) (string, string)

func drawTable(rows int, getRow getRowFn) {

}
