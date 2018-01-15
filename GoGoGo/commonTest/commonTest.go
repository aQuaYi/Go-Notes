// commonTest.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(WrapErr(WrapErr(WrapErr(errors.New("Hello, World."), "foobar"), "1"), "2"))
	fmt.Printf("%v\n", WrapErr(errors.New("my error"), "barbar"))
}
