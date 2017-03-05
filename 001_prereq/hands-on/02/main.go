package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	LicenceToKill bool
}

func (p person) pSpeak() {
	fmt.Println(p.firstName, p.lastName, fmt.Sprintf(`says : "Hello my name is %v %v!"`, p.firstName, p.lastName))
}

func (s secretAgent) sSpeak() {
	fmt.Println(s.firstName, s.lastName, fmt.Sprintf(`says : "Hello, my name is %v, %v %v!"`, s.lastName, s.firstName, s.lastName))
}

func main() {
	ps := person{"Mary", "Watson"}
	sa := secretAgent{person{"James", "Bond"}, true}

	fmt.Println(ps.firstName)
	ps.pSpeak()

	fmt.Println(sa.firstName)
	sa.sSpeak()
	sa.pSpeak()

}

//## HANDS ON 2
//- create a struct that holds person fields
//- create a struct that holds secret agent fields and embeds person type
//- attach a method to person: pSpeak
//- attach a method to secret agent: saSpeak
//- create a variable of type person
//- create a variable of type secret agent
//- print a field from person
//- run pSpeak attached to the variable of type person
//- print a field from secret agent
//- run saSpeak attached to the variable of type secret agent
//- run pSpeak attached to the variable of type secret agent
//- [SOLUTION](https://play.golang.org/p/RxrkCJw9Cd)
