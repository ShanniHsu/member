package walletFacade

import "fmt"

// Complex subsystem parts

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

// 貸方
func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

// 借方
func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}
