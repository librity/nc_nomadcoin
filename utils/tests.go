package utils

import "testing"

func Test(t *testing.T) {

	t.Run("", func(t *testing.T) {

	})

}

func ShouldPanic(t *testing.T, panicFunc func()) {
	defer func() { recover() }()
	panicFunc()

	t.Errorf("Should have panicked on function call.")
}
