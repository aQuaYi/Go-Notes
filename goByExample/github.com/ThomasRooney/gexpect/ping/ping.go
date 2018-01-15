package main

import "github.com/ThomasRooney/gexpect"
import "log"

import "strings"

func main() {
	log.Printf("Testing Ping interact... \n")

	child, err := gexpect.Spawn("ping -c8 127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer child.Close()

	for {
		l, err := child.ReadLine()
		if err != nil {
			break
		}
		if strings.Contains(l, "seq=6") {
			log.Println("进行到第6个了")
		}
		log.Println(l)
	}

	child.Interact()
	log.Printf("Success\n")
}
