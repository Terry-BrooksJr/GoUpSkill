package main 

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

type  ContactInfo struct {
	Name string
	Address string

}


func structToJSON(contact ContactInfo ) []byte {
	jsonofied, _ := json.Marshal(contact)
	return jsonofied
	}

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Contact Name")
	name, _ := reader.ReadString('\n')
	fmt.Println("Please Enter Street Address")
	address, _ := reader.ReadString('\n')

	contact := ContactInfo {
		Name: name,
		Address: address,
	}
	fmt.Println(string(structToJSON(contact)));

}