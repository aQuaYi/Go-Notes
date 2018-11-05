package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(2)
	c := big.NewInt(4)
	d := big.NewInt(8)

	d.Add(a, b).Mul(d, c)

	fmt.Println(d)

	d = c.Add(a, b)
	fmt.Println(c, d)
}

/* Output:
12
3 3
*/
