package main

import "fmt"

func main() {
	fmt.Println("Example of Maps in GoLang")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["go"] = "GoLang"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println(len(languages))

	for key, value := range languages {
		fmt.Println("Key is", key, "Value is", value)
	}
}
