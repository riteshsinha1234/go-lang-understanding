package main

import "fmt"

// power := "Golang"//shorthand cannot be declare outside the function
const age = 30

func main() {
	const name = "Golang"
	// name = "JavaScript"

	//Multiple constant

	const (
		port = 3000
		host = "localhost"
	)

	
	fmt.Println(name)
	fmt.Println(age)

	fmt.Println(port,host)
}
