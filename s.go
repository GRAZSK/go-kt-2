package main

import "fmt"

// Стратегия
type PaymentStrategy interface {
	Pay(amount int)
}

// Конкретные стратегии
type CreditCard struct{ cardNumber string }
func (c *CreditCard) Pay(amount int) {
	fmt.Printf("Paid %d using Credit Card %s\n", amount, c.cardNumber)
}

type PayPal struct{ email string }
func (p *PayPal) Pay(amount int) {
	fmt.Printf("Paid %d using PayPal account %s\n", amount, p.email)
}

// Контекст
type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

func (s *ShoppingCart) SetPaymentStrategy(ps PaymentStrategy) {
	s.paymentStrategy = ps
}

func (s *ShoppingCart) Checkout(amount int) {
	if s.paymentStrategy == nil {
		fmt.Println("Please select a payment method")
		return
	}
	s.paymentStrategy.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}
	
	cart.SetPaymentStrategy(&CreditCard{cardNumber: "1234-5678"})
	cart.Checkout(100)

	cart.SetPaymentStrategy(&PayPal{email: "test@example.com"})
	cart.Checkout(250)
}

//https://refactoring.guru/ru/design-patterns/strategy