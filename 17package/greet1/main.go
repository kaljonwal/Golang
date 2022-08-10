package main

import (
	"fmt"
	"kalyan/greet"
)

func main() {
	greet.DoGreet()

	intro := greet.Intro{
		Name: "kalyan",
		Age:  23,
	}

	fmt.Printf("%#v", intro)
}
