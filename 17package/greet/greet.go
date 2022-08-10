package greet

import "fmt"

type Intro struct {
	Name string
	Age  int
}

func DoGreet() {
	fmt.Println("DoGreet invoked")
}
