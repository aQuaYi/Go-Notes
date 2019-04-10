package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	a := 10000.
	b := a
	ratio := 0.03
	for i := 0; i < 500; i++ {
		oldA, oldB := a, b
		r := rand.Intn(10) - 4
		p := float64(100+r) / 100

		if r <= 0 {
			a, b = oldA*p, oldB*p
		} else {
			ya, yb := oldA*p-oldA, oldB*p-oldB
			fee := yb * ratio
			a, b = oldA+ya+fee, oldB+yb-fee
		}

		fmt.Printf("%10.2f, %10.2f, %10.5f, \n", a, b, a/(a+b))
	}
}
