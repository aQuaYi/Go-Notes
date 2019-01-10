package main

import (
	"log"
	"time"
)

/*
当缓冲通道投放数据时，接收端按照申请顺序依次接受。
*/

func main() {
	log.SetFlags(log.Lmicroseconds)

	bufChan := make(chan int, 3)

	for i := 0; i < 5; i++ {
		go func(id int, bufChan chan int) {
			log.Printf("%d is ready\n", id)
			log.Printf("%d: %d\n", id, <-bufChan)
		}(i, bufChan)
		time.Sleep(time.Millisecond * 200)
	}

	log.Println("All is ready...")

	for i := 0; i < 5; i++ {
		bufChan <- i
	}

	time.Sleep(time.Second)
}
