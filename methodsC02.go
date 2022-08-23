
package main

import (
  "fmt"
)

type Virtmach struct {
  ip string
  hostname string
  diskgb  int
  ram     int
}

func (v *Virtmach) show() {
  fmt.Printf("%+v\n", v)
}

func (v *Virtmach) increaseRam(add int) {
  v.ram += add
}


func main() {

  virt1 := Virtmach{"192.128.0.1", "student", 10, 5}

  virt1.show()

  virt1p := &virt1
  virt1p.increaseRam(3)

  virt1.show()
  virt1p.show()
}

