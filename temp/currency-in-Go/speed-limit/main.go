package main

import (
	"context"
	"log"
	"os"
	"sync"
)

func main() {
	defer log.Print("Done.")

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	api := open()

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			api.readFile(context.Background())
			log.Printf("%d read file", id)
		}(i)
	}

	for i := 10; i < 20; i++ {
		go func(id int) {
			defer wg.Done()
			api.resolveAddress(context.Background())
			log.Printf("%d resolve address", id)
		}(i)
	}

	wg.Wait()
}

func open() *API {
	return &API{}
}

// API is Application Programming Interface
type API struct{}

func (a *API) readFile(ctx context.Context) error {
	// do something
	return nil
}

func (a *API) resolveAddress(ctx context.Context) error {
	// do something
	return nil
}
