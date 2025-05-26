package main

import "fmt"

//
func main()  {
	fmt.Println("To find the Range")
	nums := []	int{6,7,8}

	for i:=range nums{
		fmt.Print(nums[i]," ")
	}
	fmt.Println(" ")
	fmt.Println("To Calculate the Sum of the range Range")

	sum := 0;
	nums2 := []	int{6,7,8}

	//this "_" is an index (instead of using this you can also use i)
	for _,nums := range nums2 {
		sum = sum+nums
	}
	fmt.Println(sum)

	fmt.Println("the example of Map in Range")

	map1 := map[string]string{"Frame":"John","lname":"doe"}

	for key,value := range map1{
		fmt.Println(key,value)
	}

	fmt.Println("Map for Single result for key only")

	for k:= range map1{
		fmt.Println(k)
	}

	fmt.Println("Use Range on String")

	//this give unicode value
	//starting byte of rune
	//255 -> 1 byte, 2 byte
	for i, c := range "chutiya"{
		fmt.Println(i,c)
	}




}