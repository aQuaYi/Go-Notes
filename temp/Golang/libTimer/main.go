package main

/*
很多网站的API有访问频率限制。比如，每分钟120次，等等。
那么在编写这些网站API的wrapper的时候，需要把访问频率的限制也考虑进去。
我以前的做法是把所有的网络访问需求，汇入一处，让那个goroutine管理全部的网络访问，在一个访问周期内，只访问目标网站一次。
现在想来，其实只需要安排一个售票员，保证相邻两次发票的间隔不小于一个访问周期。访问网站前，先取票，即可。
这个方法，基本可以保证先索票的goroutine，可以先得到票。
*/
import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aQuaYi/Show-You-the-Go-Code/Golang/libTimer/tickPrint"
)

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func main() {
	//产生goroutine的个数
	checkNumber := 20

	wg := &sync.WaitGroup{}
	wg.Add(checkNumber * 2)

	time.Sleep(time.Second * 3)
	for i := 0; i < checkNumber; i++ {
		if i == checkNumber /2 {
			time.Sleep(time.Second * 3 )
		}
		go tickPrint.X(i, wg)
		go tickPrint.Y(i, wg)
	}

	wg.Wait()
	fmt.Println("The End")
}
