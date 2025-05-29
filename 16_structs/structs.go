package main

import (
	"fmt"
	"time"
)

//Order Struct

type order struct{
	id string
	amount float32
	status string
	createAt time.Time //nanosecond precision
}

func (o *order) changeStatus(status string)  {
	o.status = status
}

func (o *order) getAmount() float32  {
	return o.amount
}

//to make constructor in go 

func FOrder(id string, amount float32, status string) *order{
	//initial Setup goes here...
	GenOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &GenOrder
}




func main()  {
	//if you don't set any filed, default value is zero value 
	genUOrder := FOrder("1", 30.50, "received")

	fmt.Println("This is genUOrder")
	fmt.Println(genUOrder.amount)

	 myorder := order{
		id: "1",
		amount: 50.00,
		status: "received",
	}

	myorder.changeStatus("confirm")
	fmt.Println(myorder)
	
	fmt.Println("This is for get amount")
	fmt.Println(myorder.getAmount())
	myorder.createAt = time.Now()
	myorder.status = "Paid"

	//get any field
	fmt.Println(myorder.status)


	myOrder2 := order{
		id:"2",
		amount: 100,
		status: "delivered",
		createAt: time.Now(),
	}

	fmt.Println("Order struct", myorder,myOrder2)
	
}