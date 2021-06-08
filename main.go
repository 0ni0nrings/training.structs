package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // contact filed assigned contactInfo struct
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "anderson",
		contact: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 94000,
		},
	}

	alex.updateName("alexy")
	alex.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { // use of pointer to change the name
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() { // created a receiver with a print function
	fmt.Printf("%+v\n", p)
}
