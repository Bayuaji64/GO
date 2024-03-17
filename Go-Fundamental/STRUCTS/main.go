package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy", "Tom")
	jim.print()

	// name := "Bill"

	// fmt.Println(&name)

	// name := "Bill"

	// fmt.Println(*&name)

}

func (p *person) updateName(newFirstName, newLastName string) {

	(*p).firstName = newFirstName
	(*p).lastName = newLastName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
