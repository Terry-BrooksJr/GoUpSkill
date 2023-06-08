package main 

import (
	"fmt"
	"encoding/json"
)

type struct Contact {
	name string
	address string
	city string
	state string
	postalCode int
}

func newContact(name string, address string, city string, state string, postalCode int) Contact {
	newContactAX := Contact {name : name}
	newContact.address = address
	newContact.city = city
	newContact.state = state
	newContact.postalCode = postalCode
	return newContact
}

func structToJSON(input struct){
	json.Marshall()
}
func main(){
	var postalCode int
	var name string
	fmt.Println("Please Enter Contact Name\n")
	fmt.Scan(&name)
	var address string
	fmt.Println("Please Enter Street Address\n")
	fmt.Scan(&address)
	var city string
	fmt.Println("Please Enter City\n")
	fmt.Scan(&city)
	var state string
	fmt.Println("Please Enter Street Address\n")
	fmt.Scan(&state)
	fmt.Println("Please Enter Street Address\n")
	fmt.Scan(&postalCode)

	newContact(name, address, city, state, postalCode)
}