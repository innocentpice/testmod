package waiters

import (
	"fmt"
	"testMod/counters"
	"testMod/customers"
)

type Waiters struct {
	Name    string
	Counter *counters.Counters
}

func NewWaiter(name string, counter *counters.Counters) *Waiters {
	waiter := Waiters{
		Name:    name,
		Counter: counter,
	}
	return &waiter
}

func (w *Waiters) SayHello() {
	fmt.Printf("[ %s ] : Hi ! I'm waiter from '%s'\n", w.Name, w.Counter.ShopName)
}

func (w *Waiters) ReceiveCustomer(customer *customers.Customer) {
	counter := w.Counter
	counter.RegisCustomer(customer)
	fmt.Printf("[ %s ] (Receive) : Welcome %s @ '%s'\n", w.Name, customer.Name, counter.ShopName)
}

func (w *Waiters) ReceiveOrder(customer *customers.Customer, productName string, price float32) {
	counter := w.Counter
	counter.MakeOrder(counters.OrderDetail{
		Owner:       customer.Name,
		ProductName: productName,
		Price:       price,
	})
}

func (w *Waiters) Billing(customer customers.Customer) {
	counter := w.Counter
	money := customer.PayMoney(counter.GetOrderSummary(customer))
	counter.MakePayment(customer, money)
}
