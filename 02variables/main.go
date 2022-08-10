package main

import "fmt"

var Name = "kalyan"

func main() {
	var UserName string = "kalyan"
	fmt.Println(UserName)
	fmt.Printf("variable is of type %T \n", UserName)

	var isLogedIn = false
	fmt.Println(isLogedIn)
	fmt.Printf("variable is of type %T \n", isLogedIn)

	var anotherVar float32
	fmt.Println(anotherVar)
	fmt.Printf("variable is of type %T \n", anotherVar)

}
