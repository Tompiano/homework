package main

import "fmt"

type animalAttribute struct {
	Name string
	Age  int
}

func (a animalAttribute) habit() string {
	animal := fmt.Sprintf("Name=[%v] Age=[%v]", a.Name, a.Age)
	return animal
}
func main() {
	var a animalAttribute
	a.Name = "小花"
	a.Age = 12
	fmt.Println(a)
}
