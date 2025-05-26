package main

import "fmt"


func sum(nums ...int)  int{
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
	
}


func main() {
	result := sum(3,4,5,6,)
	nums := []int{3,4,5,6}
	newResult := sum(nums...)
	
	fmt.Println(result)
	fmt.Println(nums)
	fmt.Println(newResult)
}