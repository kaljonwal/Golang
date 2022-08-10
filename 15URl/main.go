package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in/learn/account/signup?signInToPay=114900&priceId=0"

func main() {
	fmt.Println("url in go")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	// //fmt.Println(result.RawPath)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Query())
	// fmt.Println(result.Host)
	// fmt.Println(result.ForceQuery)

	buildUrl := &url.URL{
		Scheme: "http",
		Host:   "lco.dev",
		Path:   "/learn",
	}

	fmt.Println(buildUrl.String())
}
