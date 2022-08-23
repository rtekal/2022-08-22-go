
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

    inventory := make([]string, 3, 5)
    inventory[0] = "mushroom"

    mario := Player{3, 1, inventory}

  mario.show()

  mario.Greenmushroom()
  mario.show()

  mario.Pickup("item")
  mario.show()

  fmt.Println(mario.CanWhistle())

  mario.Stage = 7

  fmt.Println(mario.CanWhistle())
}
