package main

import (
	"context"
	"testing"
	"time"
)

func TestWireshark(t *testing.T) {
	//decrypTable := getConstructionTable()
	wireshark()
	// limited to uppercase
}

func TestTcpConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	serverUp := make(chan string, 3)
	defer close(serverUp)
	go demultiplexTcpReceiver(serverUp)

	for i := 0; i < 3; i++ {
		port := <-serverUp
		println("up one server ", port)
	}

	multiplexTcpSender()
	<-ctx.Done()
}

func TestRdt2_0(t *testing.T) {
	
	go receiver2_0()
	time.Sleep(2*time.Second)
	go sender2_0()
	time.Sleep(100 * time.Millisecond)
}
