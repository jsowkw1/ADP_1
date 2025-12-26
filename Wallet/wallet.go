package Wallet

import "fmt"

type Wallet struct {
	Balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.Balance += amount
	fmt.Println("Money added.")
}

func (w *Wallet) Withdraw(amount float64) {
	if amount <= w.Balance {
		w.Balance -= amount
		fmt.Println("Money withdrawn.")
	} else {
		fmt.Println("Not enough balance.")
	}
}

func (w Wallet) Info() {
	fmt.Println("Wallet balance:", w.Balance)
}
