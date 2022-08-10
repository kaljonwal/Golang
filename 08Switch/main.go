package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case statment in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := 1 //:= rand.Intn(6) + 1
	fmt.Println("Dice Number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
		fallthrough
	default:
		fmt.Println("Default case")
	}

}
