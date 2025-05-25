package main

import (
	"fmt"
	"slices"
)

//slice -> dynamic arrays
//most used construct in go
//+ useful methods
func main()  {
	//uninitialized slice is nil 
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums==nil)
	fmt.Println(len(nums))

	fmt.Println("if Slices not == nil")
	var nums2 = make([]int,2,5)
	fmt.Println(nums2)
	fmt.Println(nums2==nil)

	fmt.Println("if Slices for capacity")
	fmt.Println("The capacity of this array:",cap(nums2))

	nums2 = append(nums2,1)
	nums2 = append(nums2,2)
	nums2 = append(nums2,3)
	nums2 = append(nums2,4)
	nums2 = append(nums2,5)
	nums2 = append(nums2,6)
	nums2 = append(nums2,7)
	nums2 = append(nums2,8)
	nums2 = append(nums2,9)
	nums2 = append(nums2,10)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))


	var nums3 = make([]int,2,5)

	nums3[0] = 3
	nums3[1] = 5
	fmt.Println(nums3)

	//copy function in Slise
	

	var nums4 = make([]int,0,5)
	nums4 = append(nums4, 2)
	var nums5 = make([]int,len(nums4))
	
	//copy function
	copy(nums5,nums4)
	fmt.Println(nums4,nums5)


	//slice operator 
	fmt.Println("Slice operator")
	var nums6 = []int{1,2,3}
	fmt.Println(nums6[1:])


	//Slice package 
	fmt.Println("Slice Package")	
	var nums7 = []int{1,2,4}
	var nums8 = []int{1,2,4,5}

	fmt.Println(slices.Equal(nums7,nums8))

	//2D Slices
	fmt.Println("2D Slices")
	var nums9 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums9)
}