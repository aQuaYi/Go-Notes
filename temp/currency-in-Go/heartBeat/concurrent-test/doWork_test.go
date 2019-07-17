package work

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_doWork_unstable(t *testing.T) {
	ast := assert.New(t)

	t.Log("由于各种延迟，这个测试可能成功也可能失败")
	done := make(chan interface{})
	defer close(done)

	ints := []int{0, 1, 2, 3, 4, 5}
	_, results := doWork(done, ints...)

	for i, expected := range ints {
		select {
		case actual := <-results:
			ast.Equal(expected, actual, "index %d: expected %v, but actual %v", i, expected, actual)
		case <-time.After(1 * time.Second):
			ast.FailNow("test time out")
		}
	}
}

func Test_doWork_stable(t *testing.T) {
	ast := assert.New(t)

	done := make(chan interface{})
	defer close(done)

	ints := []int{0, 1, 2, 3, 4, 5}
	heartbeat, results := doWork(done, ints...)

	<-heartbeat // 利用 heart beat 做一次同步，避免了由于延迟带来的测试失败。

	for i, expected := range ints {
		select {
		case actual := <-results:
			ast.Equal(expected, actual, "index %d: expected %v, but actual %v", i, expected, actual)
		case <-time.After(1 * time.Second):
			ast.FailNow("test time out")
		}
	}
}
