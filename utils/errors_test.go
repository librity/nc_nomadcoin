package utils

import (
	"errors"
	"testing"
)

func TestPanicError(t *testing.T) {

	t.Run("Calls panic function when error isn't nil", func(t *testing.T) {
		oldPanicFn := panicFn
		defer func() {
			panicFn = oldPanicFn
		}()
		called := false
		panicFn = func(i ...interface{}) {
			called = true
		}

		err := errors.New("test")
		PanicError(err)

		if !called {
			t.Error("Should call panic function")
		}
	})

	t.Run("Doesn't call panic function when error is nil", func(t *testing.T) {
		oldPanicFn := panicFn
		defer func() {
			panicFn = oldPanicFn
		}()
		called := false
		panicFn = func(i ...interface{}) {
			called = true
		}

		PanicError(nil)

		if called {
			t.Error("Shouldn't call panic function")
		}
	})
}
