package main

import "fmt"

type cust struct {
	id      int
	fname   string
	lname   string
	email   string
	address string
	city    string
	state   string
	zip     int
}

func main() {

	c := cust{
		id:      0001,
		fname:   "John",
		lname:   "Doe",
		email:   "jdoe@company.xxx",
		address: "555 Main St",
		city:    "Anytown",
		state:   "California",
		zip:     11111}

	fmt.Println(c.id)
	fmt.Println(c.fname)
	fmt.Println(c.lname)
	fmt.Println(c.email)
	fmt.Println(c.address)
	fmt.Println(c.city)
	fmt.Println(c.state)
	fmt.Println(c.zip)
}
