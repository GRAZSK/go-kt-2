package main

import "fmt"

// Сложная подсистема
type Wallet struct{ balance int }

func (w *Wallet) Credit(amount int) { w.balance += amount }

type Ledger struct{}

func (l *Ledger) MakeEntry(account string, amount int) {
	fmt.Printf("Ledger entry: %s, amount: %d\n", account, amount)
}

type Notification struct{}

func (n *Notification) Send(msg string) { fmt.Println("уведомление:", msg) }

// Фасад
type WalletFacade struct {
	wallet       *Wallet
	ledger       *Ledger
	notification *Notification
}

func NewWalletFacade() *WalletFacade {
	return &WalletFacade{&Wallet{}, &Ledger{}, &Notification{}}
}

func (f *WalletFacade) AddMoney(amount int) {
	f.wallet.Credit(amount)
	f.ledger.MakeEntry("в наличие", amount)
	f.notification.Send("деньги успешно добавлены")
}

func main() {
	facade := NewWalletFacade()
	facade.AddMoney(100) // Клиент работает только с фасадом
}
//https://refactoring.guru/ru/design-patterns/facade
