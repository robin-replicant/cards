package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

// defining the custom type that is a struct
type person struct {
	fisrtName string
	lastName  string
	contact   contactInfo
}

// another way
/*
type person struct {
	fisrtName string
	lastName  string
	contactInfo // contactInfo: contactInfo
}*/

func main() {

	// go assumes the order definition of fields
	// one way
	//alex := person{"Alex", "Anderson"}

	// other way of creating a struct person
	// we can change the order of fields and it is pretty
	//alex := person{fisrtName: "Alex", lastName: "Aderson"}
	//fmt.Println(alex)

	// third way of declaring a struct
	// for structs the zero value of string is ""
	// for bool is false
	// for float of int is 0
	/*var alex person

	alex.fisrtName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	// interpolation sintax
	fmt.Printf("%+v", alex)*/

	// we must put a , even if not property comes
	jim := person{
		fisrtName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	/*
		pointer operators
		&variable give me the memory address of the value this
		variable is pointing at
		*pointer give me the value this memory address is pointing at

	*/
	// pointers option 1
	//jimPointer := &jim // reference to struct
	//jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

// pointers option 1
// we use * to pass the reference.
// *person = This is a type description, we are working with a pointer to a person
// *pointerToPerson = this is an operator, we want to manipulate the value the pointer is referencing
// Turn address into value with *address
// Turn value into address with &value
/*func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).fisrtName = newFirstName
}*/

// pointers option 2
/*
pointer shortcut:
If we define a receiver with a type of pointer
to whatever like pointer of ____, any type you can possibly
imagine.
When we attempt to call this function or we attempt to
call his method right here, go will allow us to
either call this function with a pointer or with a value
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).fisrtName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
