package main

import "fmt"

type Astro struct {
  name string
  age  int
  mission string
  isneeded bool
}

func main() {

  var astro [3]Astro

  astro[0] = Astro{"Opra", 58, "To Mars", false}
  astro[1] = Astro{"Perry", 19, "To Moon", true}
  astro[2] = Astro{"Ira", 24, "To Space", true}

  fmt.Println(astro)
  fmt.Printf("%+v\n",  astro)

}
