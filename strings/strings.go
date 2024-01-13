package main

// https://www.digitalocean.com/community/tutorials/an-introduction-to-the-strings-package-in-go

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string
	var str2 string

	str1 = "kalyan"
	str2 = "kalyan1111"

	fmt.Println("are str1 and str2 are equals :=", strings.Compare(str1, str2))

	fmt.Println("conver to uppercase", strings.ToUpper(str1))

	str3 := "kalyan jonwal"

	fmt.Println("checks whether strings starts with kalyan jo ", strings.HasPrefix(str3, "kalyan jo"))

	fmt.Println("Check whether strings ends with al", strings.HasSuffix(str3, "al"))

	fmt.Println("check if string contains an or nor", strings.Contains(str3, "an"))

	fmt.Println("Check how many times \"a\" appeared in string", strings.Count(str3, "a"))

	fmt.Println("Length of \"kalyan jonwal\" is ", len(str3))

	//The strings.Join function is useful for combining a slice of strings into a new single string.
	// To create a comma-separated string from a slice of strings, we would use this function as per the following:
	fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ","))

	//SPlit string eac it will return string array
	str := strings.Split(str3, "")

	fmt.Printf("%T\n", str)

	fmt.Println(str)

	fmt.Println(strings.Join(str, ""))

	fmt.Println(strings.ReplaceAll(str3, "jonwal", "rajput"))

}
