package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // time.Time is a struct from time package - nanosecond precision
	customer  Customer
}

func (o *Order) changeStatus(status string) {
	o.status = status
}
func (o Order) getAmount() float32 {
	return o.amount
}

func (c Customer) getCustomerPhone() string {
	return c.phone
}

func main() {
	// order := Order{
	// 	id:     "12345",
	// 	amount: 250.75,
	// 	status: "Processing",
	// }

	// order.createdAt = time.Now()
	// order.changeStatus("Shipped")

	// fmt.Println("Order ID:", order)
	// fmt.Println("Amount:", order.getAmount())

	newCustomer := Customer{
		name:  "Alice Smith",
		phone: "987-6543",
	}

	newOrder := Order{
		id:        "67890",
		amount:    150.50,
		status:    "Pending",
		customer:  newCustomer,
		createdAt: time.Now(),
	}

	newOrder.id = "54321"
	newOrder.changeStatus("Completed")

	fmt.Println("New Order ID:", newOrder)
	fmt.Println("Customer Phone:", newOrder.customer.getCustomerPhone())
}
