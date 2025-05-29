package main

import "fmt"

//Enumerated tupes
type OrderStatus string

const (
	Received OrderStatus = "received"
	Confirmed = "confirmed"
	Prepared = "Prepard"
	Delivered= "Delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status to",status)
}

func main()  {
	changeOrderStatus(Delivered)
}