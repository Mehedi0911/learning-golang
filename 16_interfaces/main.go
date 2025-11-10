package main

import "fmt"

type paymenter interface {
	pay(amount float32) // method
	refund(orderid string, amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayProcessor := razorpay{}
	// stripeProcessor := stripe{}
	// razorpayProcessor.pay(amount)
	p.gateway.pay(amount)
	fmt.Println("Payment of amount processed successfully")
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Paid using Razorpay:", amount)
}
func (r razorpay) refund(orderid string, amount float32) {
	fmt.Println("Refund using Razorpay:", amount)
}

type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Println("Paid using Stripe:", amount)
}

type fakePayment struct{}

func (r fakePayment) pay(amount float32) {
	fmt.Println("Paid using Fake Payment:", amount)
}

func main() {
	// stripeProcessor := stripe{}
	razorpayProcessor := razorpay{}
	newPayment := payment{
		gateway: razorpayProcessor,
	}
	newPayment.makePayment(300.75)
	newPayment.gateway.refund("123", 300.75)

}
