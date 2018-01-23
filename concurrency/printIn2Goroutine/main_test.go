package printIn2Goroutine

import (
	"testing"
)

func Benchmark_waitGroupPrinter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waitGroupPrinter()
	}
}

func Benchmark_channelPrinter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channelPrinter()
	}
}
