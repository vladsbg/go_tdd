package _06_ponteiros_erros

import "testing"

func TestWallet(t *testing.T) {
	checkBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		got := w.Balance()

		if want != got {
			t.Errorf(" obteve '%s', queria '%s'", got, want)
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})
}
