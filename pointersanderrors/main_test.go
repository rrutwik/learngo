package pointersanderrors

import (
	"fmt"
	"testing"
)

// TestWallet creates a Wallet, adds 10 to it using Deposit, then checks
// that the value returned from Balance is 10.
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
