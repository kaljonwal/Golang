package main

import "fmt"

type SubMersible interface {
	Dive()
}

type Shark struct {
	Name string
}

type SharkTank struct {
	Name string
}

func (s *Shark) String() string {
	return "kalyan jonwal " + s.Name
}

func (s SharkTank) String() string {
	return "kalyan jonwal " + s.Name
}

func (s *Shark) Dive() {
	s.Name = "LALLLALALA"
}

func (s *SharkTank) Dive() {
	s.Name = "jjajajaja"
}

func DeepDive(s *Shark) {
	s.Dive()
}
func main() {
	s := &Shark{
		Name: "Kalyan",
	}

	DeepDive(s)
	s1 := &SharkTank{
		Name: "Jonwal",
	}

	fmt.Println(s)
	s1.Dive()
	fmt.Println(s1)
}
