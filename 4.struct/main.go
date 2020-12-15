package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// or
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Ivo", "Ivic"}
	ivo := person{firstName: "Ivo", lastName: "Ivic"}
	fmt.Println("person ivo:", ivo)

	var marko person
	fmt.Println("person marko:", marko)

	marko.firstName = "Marko"
	marko.lastName = "Markic"
	// PrintF - print with column name value structure
	fmt.Printf("person marko updated: %+v \n", marko)

	luka := person{
		firstName: "Luka",
		lastName:  "Lukic",
		contact: contactInfo{
			email:   "luka.lukic@gmail.com",
			zipCode: 10000,
		},
	}
	fmt.Printf("person luka: %+v \n", luka)

	jozo := person{
		firstName: "Jozo",
		lastName:  "Jozic",
		contactInfo: contactInfo{
			email:   "jozo.jozic@gmail.com",
			zipCode: 12000,
		},
	}

	fmt.Printf("person jozo: %+v \n", jozo)

	jozo.print()
	jozo.updateName("Kreso")
	jozo.print()
	// & gives the memory address of the value this variable is pointing
	jozoPointer := &jozo
	jozoPointer.updateNamePointerFunction("Kreso")
	jozo.print()
	// type to pointer type (go auto allowing type or type pointer)
	jozo.updateNamePointerFunction("JozoJozo")
	jozo.print()

}

func (p person) print() {

	fmt.Printf("person print via func: %+v \n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// gives the description of the type - pointer to type
func (pointerToPerson *person) updateNamePointerFunction(newFirstName string) {
	// * gives the value of the pointing memory address
	(*pointerToPerson).firstName = newFirstName
}
