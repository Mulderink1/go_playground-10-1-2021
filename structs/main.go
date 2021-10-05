package main

import "fmt"

type person struct {
	fistName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// alex := person{fistName: "Alex", lastName: "Anderson"}

	// alex.fistName = "Mikey"

	// fmt.Println(alex)

	// var alex2 person
	// fmt.Println(alex2)
	// fmt.Printf("%+v", alex2)

	jim := person{
		fistName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "Brah",
			zipCode: 90048,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Pointer_Mikey")

	// fmt.Printf("%+v", jim)
	// jim.updateName("Mikey")
	jim.print()
}

func (p *person) updateName(newFistName string) {
	// p.fistName = newFistName
	(*p).fistName = newFistName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}