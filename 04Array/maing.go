package main

import "fmt"

func main() {
	var arr [5]string
	arr[1] = "kalyan"
	arr[2] = "Jonwal"
	fmt.Println(arr)
	fmt.Println(len(arr))

	var arr1 = [4]string{"Kalyan", "Jonwal"}

	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Printf("Type is %T", arr1)

}
