package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func getLanguages()  (string,string,string){
	return "golang","Java","python"
}



func processIt() func(a int)int  {
	return func(a int) int {
		return 4
	}	
}

func main() {
	result := add(3, 5)
	fmt.Println(result)
	fmt.Println("Print multiple values")
	fmt.Println(getLanguages())

	fmt.Println("Print multiple values using get")
	lang1,lang2,_:= getLanguages()
	fmt.Println(lang1,lang2,)

	fmt.Println("To see first class citizen in go ")
	
	fn:=processIt()
	
	fn(6)


	
}