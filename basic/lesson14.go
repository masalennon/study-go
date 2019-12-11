package main

import (
    "fmt"
)

type kelvin float64
var k kelvin = 294

type sensor func() kelvin 

func realSensor() kelvin {
    return 0
}

func calibrate(s sensor, offset kelvin) sensor {
    return func() kelvin { //無名関数を宣言することによって、型の不一致によるコンパイルエラーを防げる
        return s() + offset
    }
}

var returnK = func() kelvin {
	return k
}


func main() {
	var offset kelvin = 5
	fmt.Println("first returnK:", returnK())
	k++
	fmt.Println("second returnK:", returnK())
    sensor := calibrate(realSensor, offset)
    fmt.Println(sensor())
}