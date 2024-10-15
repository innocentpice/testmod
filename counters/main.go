package counters

import (
	"fmt"
	"testMod/customers"
	"testMod/wallets"
)

type Counters struct {
	ShopName     string
	CustomerList []*customers.Customer
	OrderList    map[string]OrderList
	Wallet       *wallets.Wallet
}

type OrderList []OrderDetail

type OrderDetail struct {
	Owner       string
	ProductName string
	Price       float32
}

func NewCounter(shopName string) *Counters {
	fmt.Printf("@Counter -> \"%s\"\n\n", shopName)
	wallet := wallets.NewWallet(0.0)
	counter := Counters{
		ShopName: shopName,
		Wallet:   &wallet,
	}
	return &counter
}

func (counter *Counters) RegisCustomer(customer *customers.Customer) {
	counter.CustomerList = append(counter.CustomerList, customer)
}

func (counter *Counters) ShowAllCustomer() {
	fmt.Println("\n------ Customer List -------")
	for index, customer := range counter.CustomerList {
		fmt.Println(index, customer)
	}
	fmt.Println("----------------------------")
}

func (counter *Counters) MakeOrder(orderDetail OrderDetail) {
	newOrderList := make(map[string]OrderList)
	for key, orderList := range counter.OrderList {
		newOrderList[key] = orderList
	}
	newOrderList[orderDetail.Owner] = append(newOrderList[orderDetail.Owner], orderDetail)
	counter.OrderList = newOrderList
}

func (counter *Counters) clearOrder(customer customers.Customer) {
	newOrderList := make(map[string]OrderList)
	for key, orderList := range counter.OrderList {
		if key != customer.Name {
			newOrderList[key] = orderList
		}
	}
	counter.OrderList = newOrderList
}

func (orders OrderList) GetSummary() float32 {
	var summary float32
	for _, order := range orders {
		summary += order.Price
	}
	return summary
}

func (counter *Counters) ShowAllOrder() {
	fmt.Println("\n--------- Order List ---------")
	for key, orderList := range counter.OrderList {
		fmt.Printf("( %s ) -> %v -> summary : %.2f\n", key, orderList, orderList.GetSummary())
	}
	fmt.Println("------------------------------")
}

func (counter *Counters) GetOrderSummary(customer customers.Customer) float32 {
	orders := counter.OrderList[customer.Name]
	return orders.GetSummary()
}

func (counter *Counters) MakePayment(customer customers.Customer, money float32) {
	counter.clearOrder(customer)
	counter.Wallet.Deposit(money)
	fmt.Printf("\n< Billing > : %s => %.2f | Wallet ( %.2f ) \n", customer.Name, money, counter.Wallet.Money)
}
