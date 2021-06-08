package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //contact filed assigned contactInfo struct
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

	fmt.Printf("%+v\n", alex)

}
