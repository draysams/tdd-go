package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		want := expected

		if got != want {
			t.Errorf("got %s but expected %s", got, want)
		}
	}

	t.Run("deposit amount", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw amount from wallet", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))

	})
}
