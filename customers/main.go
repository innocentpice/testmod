package customers

import (
	"fmt"
	"testMod/wallets"
)

type Customer struct {
	Name   string
	Wallet *wallets.Wallet
}

func NewCustomer(name string) *Customer {
	wallet := wallets.NewWallet(500.0)
	customer := Customer{
		Name:   name,
		Wallet: &wallet,
	}
	return &customer
}

func (c Customer) SayHello() {
	fmt.Printf("[ %s ] : Hi ! I'm customer\n", c.Name)
}

func (c *Customer) PayMoney(number float32) float32 {
	wallet := c.Wallet
	return wallet.WithDraw(number)
}
