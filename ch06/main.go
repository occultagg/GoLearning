package main

import "fmt"

func main() {
	type address struct {
		province string
		city string
	}

	type person struct {
		name string
		age uint

		addr address
	}

	p := person{
		age: 18,
		name: "peter",
		addr:address{
			province: "foshan",
			city: "shunde",
		},
	}

	type Stringer interface {
		String() string
	}

	func (p person) String() string {
		return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
	}

	func printString(s fmt.Stringer) {
		fmt.Println(s.String())
	}


}
