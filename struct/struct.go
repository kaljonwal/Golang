package main

import "fmt"

type employee struct {
	eid   int
	name  string
	age   int
	cname string
}

func (e1 *employee) DisplayName() {
	fmt.Println(e1.name)
	e1.name = "kalyan"
}

func (e1 *employee) DisplayAge() {
	fmt.Println(e1.age)
	e1.age = 24
}

func main() {
	e1 := &employee{
		eid:   1,
		name:  "priya",
		age:   25,
		cname: "chhota jtm",
	}

	e1.DisplayName()
	e2 := &employee{
		age: 25,
	}

	e2.DisplayAge()

	fmt.Println(e1)
	fmt.Println(e2)

	var e3 []*employee
	e3 = append(e3, e1)
	e3 = append(e3, e2)

	for _, v := range e3 {
		fmt.Println("kaka", v.age)
	}

	fmt.Println(e3)
}
