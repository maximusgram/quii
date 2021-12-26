package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
    
    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        
        wallet.Deposit(Bitcoin(10))
        
        assertBalance(t, wallet, Bitcoin(10))
    })
    
    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(100)}
        wallet.Withdraw(Bitcoin(10))
        
        assertBalance(t, wallet, Bitcoin(90))
    })
    
    t.Run("Withdraw with insufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{balance: startingBalance}
        err := wallet.Withdraw(Bitcoin(100))
        
        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()
    if got == nil {
        t.Fatal("wanted an error but didn't get one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
