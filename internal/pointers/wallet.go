package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrorInsufficientBalance = errors.New("insufficient balance")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrorInsufficientBalance
	}

	w.balance -= value

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
