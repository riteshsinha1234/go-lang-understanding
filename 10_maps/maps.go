package main

import (
	"fmt"
	"maps"
)

// map-> hash, object, dict
func main()  {
	//creating map

	m:= make(map[string]string)

	//Setting an element
	m["name"]= "golang"
	m["area"]="backend"

	//Get an element 
	fmt.Println(m["name"],m["area"])
	//IMP: if key does not exists in the map then it return zero value
	fmt.Println(m["Phone"])

	fmt.Println("Make map for different key and value")

	map1:= make(map[string]int)
	map1["age"] = 30
	map1["price"] = 50
	fmt.Println(map1["Phone"])
	fmt.Println(len(map1))
	delete(map1,"price")
	fmt.Println(map1)
	clear(map1)
	fmt.Println(map1)
	
	fmt.Println("When you have already elements available in map")
	map2 := map[string]int{"price":40,"phones":3}

	fmt.Println(map2)
	  
	fmt.Println("How to check map value on map or (Search)")

	map3 :=map[string]int{"price":40,"phones":3}

	v, Ok:= map3["phones"]

	fmt.Println(v)

	if Ok{
		fmt.Println("all ok")

	}else{
		fmt.Println("not ok")
	}


	map4 :=map[string]int{"price":40,"phones":3}
	map5 :=map[string]int{"price":40,"phones":5}

	fmt.Println(maps.Equal(map4,map5))

}