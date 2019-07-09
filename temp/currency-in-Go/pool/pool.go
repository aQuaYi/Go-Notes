package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	reuseObject()

	dynamicReuse()
}

func reuseObject() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
	// 运行了 3 次 .Get ，但是，只会创建两个对象。
}

func dynamicReuse() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	for i := 0; i < 4; i++ {
		calcPool.Put(calcPool.New())
	}

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			// 工作时间不同，
			// 创建的 Calculator 数目不同
			time.Sleep(1000 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}

func connetToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

const localAddress = "localhost:8081"

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", localAddress)
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		connPool := warmServiceConnCache()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connetToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}
