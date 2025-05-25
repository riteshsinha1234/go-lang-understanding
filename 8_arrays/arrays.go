package main

import "fmt"

//number Sequence of specific length
func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println("This is index no." ,nums[0])
	//array lenght 
	fmt.Println(len(nums))
	fmt.Println("This is full array values:- ", nums)

	//this is booleans array
	fmt.Println("This is boolean arrays")
	var vals [4]bool
	vals[2]= true
	fmt.Println(vals)

	fmt.Println("This is String arrays")

	var name[4]string
	name[0] = "golang"
	fmt.Println(name)


	//to declare array in single line
	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums2)

	//2d arrays 
	nums3 := [2][2]int{{2,4},{5,6}}
	fmt.Println(nums3)

	//- fixed size, that is predictable 
	//- memory optimazation
	//- Constant time access

	
}