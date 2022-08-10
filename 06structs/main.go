package main

import "fmt"

func main() {
	fmt.Println("Example of Structure")

	kalyan := User{
		Name:   "Kalyan",
		Email:  "kalyanjonwal555@gmail.com",
		Mobile: "+918888202879",
		Age:    21,
	}

	fmt.Println(kalyan)
	fmt.Printf("%+v", kalyan)
}

type User struct {
	Name   string
	Email  string
	Mobile string
	Age    int
}
