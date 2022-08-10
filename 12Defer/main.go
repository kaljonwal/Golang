package main

import "fmt"

func main() {
	fmt.Println("Defer Example")
	defer add(4, 2)
	defer subAndMul(4, 2)

}

func add(a int, b int) int {
	fmt.Println("Addition Function")
	return a + b
}

func subAndMul(a int, b int) (int, int) {
	fmt.Println("substraction and multiplication")
	return a - b, a * b
}
