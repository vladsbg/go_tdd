package ponteiros_erros

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.Withdraw(Bitcoin(10))
		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Retirar - saldo indisponível", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))
		confirmError(t, err, ErrorBalanceInsufficient)
	})
}

func confirmBalance(t *testing.T, w Wallet, want Bitcoin) {
	got := w.Balance()

	if want != got {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}
}

func confirmError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu")
	}

	if want != got {
		t.Errorf(" obteve '%s', queria '%s'", got, want)
	}
}
