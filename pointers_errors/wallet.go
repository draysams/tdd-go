package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String()
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("in balance address of my balance %p \n", &w.balance)
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("in deposit address of my balance %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrorInsufficientFunds
	}
	w.balance -= amount
	return nil
}
