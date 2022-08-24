package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	var myslice = []string{
		"RajTekal", "ZachFeeser", "AndrewRay", "BrunoAmorim",
		"ChrisLofgren", "KalapalaSwarupa", "KevinMcAllorum",
		"LeszekLaskowski", "Michael", "MikeH", "PatenBhavnesh",
		"PawanKumar", "Pulkit", "WilliamBillKlein", "ZaVang",
	}

	fmt.Println(myslice)

	rand.Seed(time.Now().Unix())

        r := rand.Intn(len(myslice))
        if l := len(myslice[r]); l < 8 {
		fmt.Println("Your name is short, size: ", l, " name: ", myslice[r])
	} else if l < 12 {
		fmt.Println("Your name is medium in length, size: ", l, " name: ", myslice[r])
	} else {
		fmt.Println("Your name is long, size: ", l, " name: ", myslice[r])
	}
}
