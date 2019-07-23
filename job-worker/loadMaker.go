package main

import (
	"math/rand"
)

func loadMaker() {
	for i := 0; i < 100; i++ {
		intStream <- rand.Int()
	}
}
