package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func Benchmark_NetworkRequest(b *testing.B) {
	for i := 1; i < b.N; i++ {
		conn, err := net.Dial("tcp", localAddress)
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
