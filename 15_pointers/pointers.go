package main

import "fmt"

//by value
func changeNum(num int) {
	num = 5
	fmt.Println("in changeNum",num)
}

func changeByNum(num *int)  {
	*num = 5

	fmt.Println("in changeByNum",*num)
	
}
func main() {

	num:= 1

	changeNum(num)
	changeByNum(&num)

	
	fmt.Println("Memory address", &num)
	fmt.Println("After change in main",num)





}