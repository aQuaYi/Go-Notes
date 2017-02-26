package main

import (
	"DesignModel/Adapter/adapter"
)

func main() {
	pb := adapter.NewPrintBanner("Hello")
	printWeakAndStrong(pb)
}

type print2 interface {
	PrintWeak()
	PrintStrong()
}

func printWeakAndStrong(p print2) {
	p.PrintWeak()
	p.PrintStrong()
}
