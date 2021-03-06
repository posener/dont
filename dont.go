// Package dont provides the towel you need to travel the galaxy.
package dont

// Panic gets two values, the second of them is an error. If the error is not nil, it panics with
// the error. Otherwise, it returns the first value.
// It is recommended only to use this function to initialize package scope variables, where in case
// of a panic, the program will crash on loading time.
//
// Example
//
//  import (
//      "os/user"
//      "github.com/posener/dont"
//  )
//
//  var user = dont.Panic(user.Current())
func Panic[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
