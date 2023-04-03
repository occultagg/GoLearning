package main

import "fmt"

type Age uint

func (age *Age) Modify() {
	*age = Age(30)
}

func (age Age) String() {
	fmt.Println("the age is", age)
}

func main() {
	age := Age(25)
	sm := Age.String
	sm(age)
}
