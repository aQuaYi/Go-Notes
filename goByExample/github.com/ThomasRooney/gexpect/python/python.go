package main

import (
	"log"

	"github.com/ThomasRooney/gexpect"
)

func main() {
	log.Printf("Starting python.. \n")

	child, err := gexpect.Spawn("python")
	if err != nil {
		panic(err)
	}
	defer child.Close()

	child.Start()
	child.Expect(">>>")

	child.SendLine(`print("Hellow World")`)

	child.Interact()

	log.Printf("Done \n")
}
