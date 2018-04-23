package server

import (
	"testing"
	"time"
)

func TestWebsocket(t *testing.T) {
	server := New("localhost:8888")
	done := make(chan struct{}, 1)
	result := make(chan error, 1)
	go (func() {
		result <- server.Run(done)
		close(result)
	})()
	time.Sleep(time.Second)
	close(done)
	<- result
}
