package main

import "fmt"

//for => only for-loop present in go
func main()  {
	// // while loop
	// i:=1
	// for i<=3{
	// 	fmt.Print(i," ")
	// 	i=i+1
	// }

	//infinite loop

	// for{
	// 	println("1")
	// }

	//classic for loop

	for i:= 0; i<=3; i++{
		fmt.Println(i)
	}

	//break and continue


	println("This one is new value")
	for i :=0; i<=3; i++{
		// break

		//continue
		if i == 2{
			continue
		}
		
		fmt.Println(i)

	}

	//in go 1.22 there is new feature
	// range 

	println("This is Range")
	for i := range 3 {
		fmt.Println(i)
	}


}