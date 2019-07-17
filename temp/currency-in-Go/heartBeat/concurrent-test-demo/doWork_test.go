package work

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_doWork(t *testing.T) {
	ast := assert.New(t)
	//
	done := make(chan interface{})
	defer close(done)
	//
	ints := []int{0, 1, 2, 3, 4, 5}
	timeout := 2 * time.Second
	heartbeat, results := doWork(done, timeout/2, ints...)
	//
	<-heartbeat
	//
	i := 0
	for {
		select {
		case <-heartbeat:
		case <-time.After(timeout):
			ast.FailNow("test time out")
		case actual, ok := <-results:
			if !ok {
				return
			}
			expected := ints[i]
			i++
			ast.Equal(expected, actual, "index %d: expected %d, actual %v", i, expected, actual)
		}
	}
}
