package wallet

import (
	"reflect"
	"testing"
)

func TestGetW(t *testing.T) {

	t.Run("New wallet is created", func(t *testing.T) {
		w = nil
		files = fakeFileLayer{
			fakeWalletExists: func() bool { return false },
		}

		testW := GetW()
		if reflect.TypeOf(testW) != reflect.TypeOf(&wallet{}) {
			t.Error("Should return a new wallet instance.")
		}
	})

	t.Run("Wallet is initialized from file", func(t *testing.T) {
		w = nil
		files = fakeFileLayer{
			fakeWalletExists: func() bool { return true },
		}

		testW := GetW()
		if reflect.TypeOf(testW) != reflect.TypeOf(&wallet{}) {
			t.Error("Should return a new wallet instance.")
		}

		dResult := testW.privateKey.D
		dExpected := testWallet.privateKey.D
		if dResult.Cmp(dExpected) != 0 {
			t.Errorf("Expected %v, got %v", dExpected, dResult)
		}
	})

}

func TestGetAddress(t *testing.T) {

	t.Run("Should return address", func(t *testing.T) {
		w = nil
		files = fakeFileLayer{
			fakeWalletExists: func() bool { return true },
		}

		expected := "c544a8f9319ad8e2688a4c66c94cb6ed434b86058ad1489d4d152ff6fc47943718ce246f35ef03c8b3dce4e39bf409a3016830a5b1eef6c441453951586b2903"
		result := GetAddress()
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

}
