package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File handling in GO")

	file, err := os.Create("./fileexample.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(*file)
	defer file.Close()

	// writing into the file
	ipToFile := "Employee: Kalyan Jonwal Employer: Aqua security"

	length, err := io.WriteString(file, ipToFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("length of file is ", length)

	readFile("./fileexample.txt")
}

func readFile(filePath string) {
	bytearr, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytearr))
}
