package main

import "fmt"

func main() {
	fmt.Println("Functions Example")

	greeting()

	adder := addreFunc(5, 3)
	fmt.Println("Result of adder is ", adder)

	proAdder := proAdderFunc(5, 5, 5, 5, 5)
	fmt.Println("Result of proadder is", proAdder)
}

func proAdderFunc(elements ...int) int {
	total := 0
	for _, value := range elements {
		total += value
	}
	return total
}
func addreFunc(ele1 int, ele2 int) int {
	return ele1 + ele2
}
func greeting() {
	fmt.Println("Namastey India")
}
