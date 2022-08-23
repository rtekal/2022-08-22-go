package main

import "fmt"


type Astro struct {
  name string
  age  int
  mission string
  isneeded bool
}

type nasaMission struct {
     people []Astro
     number int
     message string
 }


func main() {

    p := make([]Astro, 3, 5)

  p[0] = Astro{"Opra", 58, "To Mars", false}
  p[1] = Astro{"Perry", 19, "To Moon", true}
  p[2] = Astro{"Ira", 24, "To Space", true}

  fmt.Println(p)
  fmt.Printf("%+v\n",  p)

  var s nasaMission
  s.people = p
  s.number = 3
  s.message = "success"

    fmt.Printf("%+v\n", s)

}
