package dont

import (
	"testing"
	"fmt"
)

func TestPanic(t *testing.T) {
    t.Parallel()
    
    want := "Yay!"
    f := func() (string, error) { return want, nil }

    if got := Panic(f()); got != want {
        t.Errorf("Panic(good()) = %v, want %v", got, want)
    }
}

func TestPanicError(t *testing.T) {
    t.Parallel()

    want := fmt.Errorf("failed")
    f := func() (string, error) { return "", want }

    defer func() {
        got := recover()
        if got != want {
            t.Errorf("Panic(bad()) error is %v, want %v", got, want)
        }
    }()

    Panic(f())
}