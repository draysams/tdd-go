package pointers_errors

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("in test address of my balance %p \n", &wallet.balance)

	balance := Bitcoin(10)

	if got != balance {
		t.Errorf("got %s but expected %s", got, balance)
	}
}
