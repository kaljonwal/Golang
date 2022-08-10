package main

import "fmt"

func main() {
	fmt.Println("if else example in golang")

	value := 10

	if value < 10 {
		fmt.Println("Value is less than 10")
	} else if value > 10 {
		fmt.Println("Value is greater than 10")
	} else {
		fmt.Println("Value is 10")
	}

	if num := 3; value > num {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
