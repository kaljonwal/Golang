package main

import "fmt"

func main() {
	fmt.Println("Example of methods in Golang")

	userDetails := User{
		UserName: "kalyan",
		Age:      22,
		Email:    "kalyanjonwal555@gmail.com",
	}

	fmt.Println(userDetails)
	userDetails.SetUser("kalyan.jonwal@aquasec.com")
	fmt.Println(userDetails)

}

type User struct {
	UserName string
	Age      int
	Email    string
}

func (u *User) SetUser(email string) {
	u.Email = email
}
