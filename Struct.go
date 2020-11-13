package main

import "fmt"

// Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"Akash", "Dehradun", 249404}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Verma", city: "Patna",
		Pincode: 800001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)
}
