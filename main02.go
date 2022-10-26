package main

import "fmt"

type Interface interface {
	together()
}

type Dog struct {
	Name  string
	Age   int
	Habit string
}

func (a *Dog) together() {
	str := fmt.Sprintf("Name=[%v] Age=[%v] habit=[%v]", a.Name, a.Age, a.Habit)
	fmt.Println(str)
}

func main() {
	var a Interface = &Dog{
		Name:  "小花",
		Age:   2,
		Habit: "汪汪",
	}
	a.together()
}
