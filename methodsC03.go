
package main

import (
  "fmt"
)

type Player struct {
  Lives int
  Stage int
  Inventory []string
}

func (p *Player) show() {
  fmt.Printf("%+v\n", p)
}

func (p *Player) Greenmushroom() {
  p.Lives++
}

func (p *Player) Pickup(item string) {
  p.Inventory = append(p.Inventory, item)
}

func (p *Player) CanWhistle() bool {
  return p.Stage >= 5
}

func main() {

  mario := Player{3, 1, make([]string, 3, 5)}

  mario.show()

  pmario := &mario
  pmario.Greenmushroom()
  pmario.show()

  pmario.Pickup("item")
  pmario.show()

  pmario.CanWhistle()

}
