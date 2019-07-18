package main

import (
	"context"
	"log"
	"os"
	"sort"
	"sync"
	"time"

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

func per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func open() *API {
	secondLimit := rate.NewLimiter(per(2, time.Second), 1)
	minuteLimit := rate.NewLimiter(per(10, time.Minute), 10)
	return &API{
		rateLimiter: makeMultiLimiter(secondLimit, minuteLimit),
	}
}

// API is Application Programming Interface
type API struct {
	rateLimiter rateLimiter
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

type rateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

type multiLimiter struct {
	limiters []rateLimiter
}

func makeMultiLimiter(limiters ...rateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{
		limiters: limiters,
	}
}

func (ml *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range ml.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ml *multiLimiter) Limit() rate.Limit {
	return ml.limiters[0].Limit()
}
