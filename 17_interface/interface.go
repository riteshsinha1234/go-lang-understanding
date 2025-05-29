package main

import "fmt"

type payment struct {
	gateway stripe
}


//
func (p payment) makePaymet(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	
	p.gateway.pay(amount)
	
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorepay",amount)

}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe",amount)
}

func main() {
	fmt.Println("This is for Razorepay")
	stripePaymentGw:= stripe{}
	newPayment:= payment{
		gateway: stripePaymentGw,
	}
	newPayment.makePaymet(100)
}