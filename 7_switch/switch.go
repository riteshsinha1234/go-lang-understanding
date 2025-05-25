package main

import (
	"fmt"
	"time"
)

func main() {
	//Simple Swich
	fmt.Println("This is multiple Switch")
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("tow")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Other")
	}

	//Multiple condition switch
	fmt.Println("This is multiple Switch")
	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("it's weekend")
		default:
			fmt.Println("it's workday")
		}
	
	//type switch 
	fmt.Println("This is Type Switch")
	whoAmI := func (i interface{})  {
		switch i.(type){
		case int:
			fmt.Println("it's a integer")
		case string:
			fmt.Println("it's a String")
		case bool:
			fmt.Println("it's a boolean")

		default:
			fmt.Println("Other")
		}


	}
	whoAmI(586.39)

}