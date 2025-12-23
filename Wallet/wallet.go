package Wallet

import "fmt"

type Wallet struct {
	Owner   string
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	fmt.Printf("%s added %.2f. New balance: %.2f\n", w.Owner, amount, w.Balance)
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount <= w.Balance {
		w.Balance -= amount
		fmt.Printf("%s spent %.2f. Remaining balance: %.2f\n", w.Owner, amount, w.Balance)
	} else {
		fmt.Println("Not enough balance")
	}
}
