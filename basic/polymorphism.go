package main

import (
  "fmt"
  "strconv"
)

func main() {
  myCar := &prius{name: "マイカー", speed: 0}
  fmt.Println(myCar)
  hisCar := &lexus{name: "彼の車", speed: 0}

  var protoA prius
  protoA.name = "proto A"
  protoA.speed = 10

  fmt.Printf("%+v", protoA)

  var c car = myCar //interface変数に入れることで、インターフェイスを満たしていることが確認できる。
  fmt.Println(c.run(30))
  c.stop()

  var sample struct{
    name string
    age int
  }

  fmt.Println(sample)

    
  priuses := []prius{
    {name: "Prius A", speed: 10},
    {name: "Prius B", speed: 10},
  }
  fmt.Printf("\n%+v \n", priuses)
  

  var h car = hisCar
  fmt.Println(h.run(40))
  h.stop()
}
type car interface {
  run(int) string
  stop()
}

type prius struct {
  name  string
  speed int
}

type lexus struct {
  name string
  speed int
}

//メソッドの実装
func (u *prius) run(speed int) string {
  u.speed = speed
  return strconv.Itoa(speed) + "kmで走ります"
}
func (u *prius) stop() {
  fmt.Println("停止します")
  u.speed = 0
}
func (l *lexus) run(speed int) string {
  l.speed = speed * 3
  return "Run at the speed of " + strconv.Itoa(speed) + "km"
}

func (l *lexus) stop() {
  fmt.Println("Stopping...")
  l.speed = 0
}