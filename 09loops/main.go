package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	sports := []string{}
	sports = append(sports, "Cricket", "Soccer", "base ball", "football")
	fmt.Println(sports)

	// for i := 0; i < len(sports); i++ {
	// 	fmt.Println(sports[i])
	// }

	// for i := range sports {
	// 	fmt.Println(sports[i])
	// }

	for index, value := range sports {
		if index == 3 {
			goto gotoloop
		}
		fmt.Println("Index is", index, " Value is", value)
	}

gotoloop:
	fmt.Println("Hits the goto")
}
