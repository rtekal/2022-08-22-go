package main

import (
    "fmt"
)

func main() {
 var (
  uri string = "https://example.org:6001/v2/snacks?"
  r   string = "req=snickers"
  q   string = "quantity=1"
  s   string = "size=king"
  )

  furi := fmt.Sprintf("%s%s&%s&%s\n",uri, r, q, s)

  fmt.Println(furi)
}
