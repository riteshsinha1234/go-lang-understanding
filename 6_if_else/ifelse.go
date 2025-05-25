package main

import "fmt"

func main() {
	fmt.Println("This is if-else")
	age := 16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	fmt.Println("This is else-if")

	if age >=18{
		fmt.Println("Person is an adult")


	}else if age>=12{
			fmt.Println("Person is an teenager")
	}else{
		fmt.Println("Person is a kid")
	}

	fmt.Println("This is Logical operator")

	var role = "admin"
	var hasPermissions = false
	
	if role == "admin" && hasPermissions{
		fmt.Println("Yes")
	}

	fmt.Println("Declare variable inside the if")

	if age := 20; age >=18{
		fmt.Println("Person is an adult",age)
	}else if age >= 12 {
		fmt.Println("Person is an teenager",age)
	}

	//go does not have ternary operator, you will have to use normal if else

}