package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/time/rate"
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
			err := api.readFile(context.Background())
			if err != nil {
				log.Printf("%d cannot read file: %s", id, err)
			}

			log.Printf("%d read file", id)
		}(i)
	}

	for i := 10; i < 20; i++ {
		go func(id int) {
			defer wg.Done()
			err := api.resolveAddress(context.Background())
			if err != nil {
				log.Printf("%d cannot resolve address: %s", id, err)
			}
			log.Printf("%d resolve address", id)
		}(i)
	}

	wg.Wait()
}

func open() *API {
	return &API{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

// API is Application Programming Interface
type API struct {
	rateLimiter *rate.Limiter
}

func (a *API) readFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// do something
	return nil
}

func (a *API) resolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// do something
	return nil
}
