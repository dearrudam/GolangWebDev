package main

import (
	"fmt"
)

type normalPerson struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	normalPerson
	LicenceToKill bool
}

func (p normalPerson) speak() {
	fmt.Println(p.firstName, p.lastName, fmt.Sprintf(`says : "Hello my name is %v %v!"`, p.firstName, p.lastName))
}

func (s secretAgent) speak() {
	fmt.Println(s.firstName, s.lastName, fmt.Sprintf(`says : "Hello, my name is %v, %v %v!"`, s.lastName, s.firstName, s.lastName))
}

type person interface {
	speak()
}

func print(p person) {
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	switch v := p.(type) {
	case normalPerson:
		fmt.Println(v.firstName)
	case secretAgent:
		fmt.Println(v.firstName)
	default:
		fmt.Println("unknown")
	}
	p.speak()
}

func main() {
	ps := normalPerson{"Mary", "Watson"}
	sa := secretAgent{
		normalPerson{
			firstName: "James",
			lastName:  "Bond",
		},
		true,
	}

	print(ps)
	print(sa)
	sa.normalPerson.speak()

}

//## HANDS ON 3
//- create an interface type that both person and secretAgent implement
//- declare a func with a parameter of the interface’s type
//- call that func in main and pass in a value of type person
//- call that func in main and pass in a value of type secretAgent
//- [SOLUTION](https://play.golang.org/p/-Ux0gHf4SF)