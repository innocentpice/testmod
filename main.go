package main

import (
	"testMod/counters"
	"testMod/customers"
	"testMod/waiters"
)

func main() {

	counter := counters.NewCounter("CG SHOP")

	customer := customers.NewCustomer("Spiceze")
	customer2 := customers.NewCustomer("Bella")

	waiter := waiters.NewWaiter("Bobby", counter)
	waiter2 := waiters.NewWaiter("Max", counter)

	customer.SayHello()
	customer2.SayHello()
	waiter.SayHello()
	waiter2.SayHello()

	println("---------------------")

	waiter.ReceiveCustomer(customer)
	// counter.ShowAllCustomer()

	println("---------------------")

	waiter2.ReceiveCustomer(customer2)
	// counter.ShowAllCustomer()

	println("---------------------")

	waiter.ReceiveOrder(customer, "Chicken", 150.0)
	waiter.ReceiveOrder(customer, "Hamberg", 220.0)
	waiter.ReceiveOrder(customer2, "Rice", 150.0)

	counter.ShowAllOrder()

	println("---------------------")

	waiter2.ReceiveOrder(customer2, "Juice", 50.0)
	waiter2.ReceiveOrder(customer, "Water", 30.0)

	counter.ShowAllOrder()

	println("---------------------")

	waiter.Billing(*customer)
	counter.ShowAllOrder()
	println("---------------------")

	waiter2.Billing(*customer2)
	counter.ShowAllOrder()

}
