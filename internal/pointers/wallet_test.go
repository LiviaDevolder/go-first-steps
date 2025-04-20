package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	verifyBalance := func(t *testing.T, wallet Wallet, expect Bitcoin) {
		t.Helper()

		// Act
		result := wallet.Balance()

		// Assert
		if result != expect {
			t.Errorf("result %s, expect %s", result, expect)
		}
	}

	verifyError := func(t *testing.T, result error, expect error) {
		t.Helper()

		if result == nil {
			t.Fatal("Expected an error but none occurred")
		}

		if result != expect {
			t.Errorf("result %s, expect %s", result, expect)
		}
	}

	verifyInexistentError := func(t *testing.T, result error) {
		t.Helper()

		if result != nil {
			t.Fatal("unexpected error")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		// Arrange
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expect := Bitcoin(10)

		// Act & Assert
		verifyBalance(t, wallet, expect)
	})

	t.Run("Withdraw", func(t *testing.T) {
		// Arrange
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		expect := Bitcoin(10)

		// Act & Assert
		verifyBalance(t, wallet, expect)
		verifyInexistentError(t, err)
	})

	t.Run("Withdraw without enough balance", func(t *testing.T) {
		// Arrange
		initialBalance := Bitcoin(20)

		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		// Act & Assert
		verifyBalance(t, wallet, initialBalance)
		verifyError(t, err, ErrorInsufficientBalance)
	})
}
