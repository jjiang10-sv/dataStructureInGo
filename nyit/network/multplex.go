package main

import (
	"fmt"
	"net"
	"sync"
)

func multiplexSender() {
	// Define the addresses of the receivers
	receivers := []string{
		"127.0.0.1:8001",
		"127.0.0.1:8002",
		"127.0.0.1:8003",
	}

	// Create a UDP connection
	conn, err := net.Dial("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("Error creating UDP connection:", err)
		return
	}
	defer conn.Close()

	// Send data to each receiver
	for _, addr := range receivers {
		message := fmt.Sprintf("Hello, receiver at %s", addr)
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
		fmt.Printf("Sent: %s to %s\n", message, addr)
	}
}

func demultiplexReceiver() {
	// Define the port to listen on
	port := ":8001" // Change this to 8002, 8003, etc., for different receivers

	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Listening on %s\n", port)

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Print the received data
		fmt.Printf("Received: %s from %s\n", string(buffer[:n]), addr)
	}
}

func multiplexTcpSender() {
	// Define server addresses (ports) to connect to
	servers := []string{
		"127.0.0.1:8001",
		"127.0.0.1:8002",
		"127.0.0.1:8003",
	}

	var wg sync.WaitGroup

	// Connect to each server and send data
	for _, addr := range servers {
		wg.Add(1)
		go func(serverAddr string) {
			defer wg.Done()

			// Establish a TCP connection
			conn, err := net.Dial("tcp", serverAddr)
			if err != nil {
				fmt.Printf("Error connecting to %s: %v\n", serverAddr, err)
				return
			}
			defer conn.Close()

			// Send data to the server
			message := fmt.Sprintf("Hello, server at %s", serverAddr)
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Printf("Error sending data to %s: %v\n", serverAddr, err)
				return
			}

			fmt.Printf("Sent: %s to %s\n", message, serverAddr)
		}(addr)
	}

	// Wait for all connections to finish
	wg.Wait()
}

func handleConnection(conn net.Conn, port string) {
	defer conn.Close()

	// Buffer to hold incoming data
	buffer := make([]byte, 1024)

	// Read data from the connection
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from %s: %v\n", port, err)
		return
	}

	// Print the received data
	fmt.Printf("Received on port %s: %s\n", port, string(buffer[:n]))
}

func startServer(port string, wg *sync.WaitGroup , serverUpSignal chan(string)) {
	defer wg.Done()

	// Create a TCP address to listen on
	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		fmt.Printf("Error resolving TCP address for %s: %v\n", port, err)
		return
	}

	// Start listening on the specified port
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("Error listening on %s: %v\n", port, err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", port)
	serverUpSignal <- port

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection on %s: %v\n", port, err)
			continue
		}

		// Handle the connection in a new goroutine
		handleConnection(conn, port)
	}
}

func demultiplexTcpReceiver(serverUp chan(string)) {
	// Define the ports to listen on
	ports := []string{
		":8001",
		":8002",
		":8003",
	}

	var wg sync.WaitGroup

	// Start a server for each port
	for _, port := range ports {
		wg.Add(1)
		go startServer(port, &wg,serverUp)
	}

	// Wait for all servers to finish (they won't, as they run indefinitely)
	wg.Wait()
}
