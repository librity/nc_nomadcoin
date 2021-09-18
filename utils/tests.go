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

func FailIfDifferent(t *testing.T, expected interface{}, result interface{}) {
	if expected != result {
		ErrorDifferent(t, expected, result)
	}
}

func ErrorDifferent(t *testing.T, expected interface{}, result interface{}) {
	t.Errorf("Expected %v, got %v", expected, result)
}
