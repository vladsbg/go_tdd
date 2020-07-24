package ponteiros_erros

import (
	"errors"
	"fmt"
)

var ErrorBalanceInsufficient = errors.New("Saldo indisponível")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrorBalanceInsufficient

	}
	w.balance -= value
	return nil
}
