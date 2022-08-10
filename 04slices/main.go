package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}
	fruitList = append(fruitList, "kalyan")
	fruitList = append(fruitList, "jonwal")
	fruitList = append(fruitList, "pict", "pune", "Aquasec", "Hyd")
	fmt.Println(fruitList)

	fruitList = append(fruitList[4:])
	fmt.Println(fruitList)

	hisghScores := make([]int, 0)
	hisghScores = append(hisghScores, 5, 2, 3, 4)
	fmt.Println(hisghScores)

	fmt.Println(sort.IntsAreSorted(hisghScores))

	//How to remove values based on inddex in slices

	var courses = []string{"c", "c++", "java", "Go", "python", "js"}
	var index int = 1
	fmt.Println("Courses before removing", courses)

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
