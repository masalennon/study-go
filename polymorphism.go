package main

import (

)

func main() {

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
func (l *lexus) run() {
  l.speed = speed * 1.2
  return strconv.Itoa(speed) + "Run at the speed of"
}

func (u *prius) stop() {
  fmt.Println("Stopping...")
  u.speed = 0
}