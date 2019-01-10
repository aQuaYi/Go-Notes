package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "say hello to whom.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s\n", name)
}
