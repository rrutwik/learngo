package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin uint

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Deposit adds the given amount to the balance of the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Deposit: address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Balance returns the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("Balance: address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}

var ErrNoMoney = errors.New("can't withdraw, insufficient funds")

// Withdraw subtracts the given amount from the balance of the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	fmt.Printf("Withdraw: address of balance in Deposit is %p \n", &w.balance)
	if amount > w.balance {
		return ErrNoMoney
	}
	w.balance -= amount
	return nil
}
