package main

import "fmt"

func main() {

  var fileExtensions = map[string]string{
    "Python": ".py",
    "Golang": ".go",
    "Java":   ".java",
    "Ansible": ".yml",
    "C++":     ".cpp",
    }

    fmt.Println(fileExtensions)

    v, ok := fileExtensions["Golang"]
    if !ok {
       fmt.Println("Not found")
    } else {
       fmt.Println("Found: ", v)
    }

    delete(fileExtensions, "C++")
    fmt.Println(fileExtensions)

    fileExtensions["Julia"] = ".jl"

    fmt.Println(fileExtensions)

}
