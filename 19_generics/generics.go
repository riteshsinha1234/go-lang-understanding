package main

import "fmt"

func printSlice[T int | string ](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements[] T
}



func main() {
	mystack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println(mystack)

	names:= []string {"golang","typescript"}
	printSlice(names)
}