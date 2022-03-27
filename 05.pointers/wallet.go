package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin)  {
    w.balance += money
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

func (w *Wallet) Withdraw(money Bitcoin) error {
    if money > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= money
    return nil
}

