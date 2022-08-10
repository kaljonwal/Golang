package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handling web request in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

	databytes, err := ioutil.ReadAll(response.Body)

	op := string(databytes)
	fmt.Println("Body is ", op)
}
