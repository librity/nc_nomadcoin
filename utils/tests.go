package utils

import "testing"

func ShouldPanic(t *testing.T, panicFunc func()) {
	defer func() { recover() }()
	panicFunc()

	t.Errorf("Should have panicked on function call.")
}
