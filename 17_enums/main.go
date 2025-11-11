package main

import "fmt"

type OrderStatus int
type OrderPaymentMethod string

const (
	StatusPending OrderStatus = iota
	StatusShipped
	StatusDelivered
	StatusCancelled
)
const (
	CreditCard   OrderPaymentMethod = "Credit Card"
	PayPal                          = "PayPal"
	BankTransfer                    = "Bank Transfer"
)

func changeStatus(s OrderStatus) {
	fmt.Println("Changing status to:", s)
}
func selectPaymentMethod(p OrderPaymentMethod) {
	fmt.Println("Changing status to:", p)
}

func main() {
	changeStatus(StatusDelivered)
	selectPaymentMethod(BankTransfer)
}
