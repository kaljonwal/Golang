package main

import "fmt"

func main() {
	//var mapds map[string]string
	mapds := make(map[string]string)
	mapds["name"] = "kalyan"
	mapds["age"] = "24"
	mapds["company"] = "Aqua"
	mapds["City"] = "Phulambri"

	var keyslice []string
	var valueslice []string

	for k, v := range mapds {
		//fmt.Println(k, "-->", v)
		keyslice = append(keyslice, k)
		valueslice = append(valueslice, v)
	}
	fmt.Println(mapds)
	fmt.Println(keyslice)
	fmt.Println(valueslice)

	//checking existence
	count, ok := mapds["name"]
	fmt.Println(count, ok)

	count1, ok := mapds["age"]
	fmt.Println(count1, ok)

}
