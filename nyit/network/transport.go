package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	receiverAddr = "localhost:9999"
	serverAddr   = "localhost:9999"

	windowSize   = 4
	totalPackets = 10
	timeout      = 2 * time.Second
	maxRetries   = 5
)

const (
	listenAddr = "localhost:9999"
	packetLoss = 20 // Simulating 20% packet loss
)

func selectiveRepeatSender() {
	conn, err := net.Dial("udp", receiverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	base := 0
	nextSeqNum := 0
	acks := make(chan int, totalPackets)
	var wg sync.WaitGroup
	sentPackets := make(map[int]bool)
	mutex := &sync.Mutex{}

	// Goroutine to listen for ACKs
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _ := conn.Read(buf)
			ack, _ := strconv.Atoi(string(buf[:n]))
			acks <- ack
		}
	}()

	for base < totalPackets {
		// Send packets within the window
		for nextSeqNum < base+windowSize && nextSeqNum < totalPackets {
			if !sentPackets[nextSeqNum] { // Only send unsent packets
				packet := fmt.Sprintf("%d|DATA-PACKET", nextSeqNum)
				fmt.Println("Sending:", packet)
				_, _ = conn.Write([]byte(packet))
				sentPackets[nextSeqNum] = true
			}
			nextSeqNum++
		}

		// Set a timer for retransmission
		timer := time.NewTimer(timeout)
		retries := 0

		select {
		case ack := <-acks:
			mutex.Lock()
			fmt.Println("Received ACK:", ack)
			delete(sentPackets, ack) // Mark packet as acknowledged

			// Move base forward to the next unacknowledged packet
			for base < totalPackets && !sentPackets[base] {
				base++
			}
			mutex.Unlock()
			timer.Stop()

		case <-timer.C:
			fmt.Println("Timeout! Resending lost packets...")
			mutex.Lock()
			for seq := range sentPackets {
				if seq >= base && seq < base+windowSize {
					packet := fmt.Sprintf("%d|DATA-PACKET", seq)
					fmt.Println("Resending:", packet)
					_, _ = conn.Write([]byte(packet))
				}
			}
			mutex.Unlock()

			retries++
			if retries >= maxRetries {
				fmt.Println("Max retries reached. Terminating.")
				return
			}
		}
	}

	wg.Wait()
	fmt.Println("All packets successfully sent and acknowledged!")
}

func selectiveRepeatReceiver() {
	addr, _ := net.ResolveUDPAddr("udp", listenAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting receiver:", err)
		return
	}
	defer conn.Close()

	expectedSeq := 0
	buffer := make(map[int]string) // Buffer for out-of-order packets
	mutex := &sync.Mutex{}
	buf := make([]byte, 1024)

	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buf)
		message := string(buf[:n])
		parts := strings.SplitN(message, "|", 2)

		if len(parts) != 2 {
			continue
		}

		seqNum, _ := strconv.Atoi(parts[0])
		data := parts[1]

		// Simulate packet loss
		if rand.Intn(100) < packetLoss {
			fmt.Println("Packet loss simulated. Dropping packet:", seqNum)
			continue
		}

		mutex.Lock()
		// Store packet in buffer
		buffer[seqNum] = data

		// Process in-order packets and shift window
		for {
			if val, exists := buffer[expectedSeq]; exists {
				fmt.Println("Delivered:", val, "Seq:", expectedSeq)
				delete(buffer, expectedSeq)
				expectedSeq++
			} else {
				break
			}
		}

		// Send ACK for received packet
		ackMsg := strconv.Itoa(seqNum)
		conn.WriteToUDP([]byte(ackMsg), remoteAddr)
		mutex.Unlock()
	}
}

func pipelineGoBackNSender() {
	conn, err := net.Dial("udp", receiverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	base := 0
	nextSeqNum := 0
	acks := make(chan int)
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}

	// Goroutine to listen for ACKs
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _ := conn.Read(buf)
			ack, _ := strconv.Atoi(string(buf[:n]))
			acks <- ack
		}
	}()

	for base < totalPackets {
		// Send packets within the window
		for nextSeqNum < base+windowSize && nextSeqNum < totalPackets {
			packet := fmt.Sprintf("%d|DATA-PACKET", nextSeqNum)
			fmt.Println("Sending:", packet)
			_, _ = conn.Write([]byte(packet))
			nextSeqNum++
		}

		// Set a timer for retransmission
		timer := time.NewTimer(timeout)
		retries := 0

		select {
		case ack := <-acks:
			mutex.Lock()
			if ack >= base {
				fmt.Println("Received ACK:", ack)
				base = ack + 1
			}
			mutex.Unlock()
			timer.Stop()

		case <-timer.C:
			fmt.Println("Timeout! Resending all unacknowledged packets...")
			retries++
			if retries >= maxRetries {
				fmt.Println("Max retries reached. Terminating.")
				return
			}
			nextSeqNum = base
		}
	}

	wg.Wait()
	fmt.Println("All packets successfully sent and acknowledged!")
}

func pipelineGoBackNReceiver() {
	addr, _ := net.ResolveUDPAddr("udp", listenAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting receiver:", err)
		return
	}
	defer conn.Close()

	expectedSeq := 0
	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)
		message := string(buffer[:n])
		parts := strings.SplitN(message, "|", 2)

		if len(parts) != 2 {
			continue
		}

		seqNum, _ := strconv.Atoi(parts[0])
		data := parts[1]

		// Simulate packet loss
		if rand.Intn(100) < packetLoss {
			fmt.Println("Packet loss simulated. Dropping packet:", seqNum)
			continue
		}

		// Process the received packet
		if seqNum == expectedSeq {
			fmt.Println("Received:", data, "Seq:", seqNum)
			expectedSeq++
		} else {
			fmt.Println("Out-of-order packet received. Ignoring:", seqNum)
		}

		// Send cumulative ACK
		ackMsg := strconv.Itoa(expectedSeq - 1)
		conn.WriteToUDP([]byte(ackMsg), remoteAddr)
	}
}

func sender3_0() {
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	messages := []string{"Hello", "Reliable", "Data", "Transfer"}
	seqNum := 0 // Alternating sequence number (0 or 1)

	for _, msg := range messages {
		packet := fmt.Sprintf("%d|%s", seqNum, msg)
		retries := 0

		for {
			// Send the packet
			fmt.Println("Sending:", packet)
			_, err = conn.Write([]byte(packet))
			if err != nil {
				fmt.Println("Send error:", err)
				return
			}

			// Start the retransmission timer
			timer := time.NewTimer(timeout)

			// Channel to signal successful ACK receipt
			ackReceived := make(chan bool)

			// Goroutine to listen for ACK
			go func() {
				buf := make([]byte, 1024)
				conn.SetReadDeadline(time.Now().Add(timeout))
				n, err := conn.Read(buf)

				if err == nil {
					ack := string(buf[:n])
					fmt.Println("Received:", ack)
					if ack == fmt.Sprintf("ACK%d", seqNum) {
						ackReceived <- true
					}
				}
			}()

			select {
			case <-ackReceived:
				// ACK received successfully, move to the next message
				timer.Stop()
				seqNum = 1 - seqNum // Toggle sequence number
				break
			case <-timer.C:
				// Timeout occurred, retransmit
				fmt.Println("Timeout! Retransmitting:", packet)
				retries++
				if retries >= maxRetries {
					fmt.Println("Max retries reached, giving up on this packet.")
					return
				}
			}
		}
	}
}

func sender2_1() {
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	messages := []string{"Hello", "Reliable", "Data", "Transfer"}
	seqNum := 0 // Alternating sequence number (0 or 1)

	for _, msg := range messages {
		packet := fmt.Sprintf("%d|%s", seqNum, msg)
		for {
			fmt.Println("Sending:", packet)
			_, err = conn.Write([]byte(packet))
			if err != nil {
				fmt.Println("Send error:", err)
				return
			}

			// Set timeout for ACK/NAK
			conn.SetReadDeadline(time.Now().Add(timeout))
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Timeout: Retransmitting", packet)
				continue
			}

			ack := string(buf[:n])
			fmt.Println("Received:", ack)

			// Check if ACK matches expected sequence number
			if ack == fmt.Sprintf("ACK%d", seqNum) {
				seqNum = 1 - seqNum // Toggle sequence number
				break
			} else {
				fmt.Println("Garbled ACK! Retransmitting", packet)
			}
		}
	}
}

func sender2_0() {
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	messages := []string{"Hello", "Reliable", "Data", "Transfer"}

	for _, msg := range messages {
		fmt.Println("Sending:", msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Send error:", err)
			return
		}

		// Set timeout for ACK/NAK
		conn.SetReadDeadline(time.Now().Add(timeout))
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Timeout: Retransmitting", msg)
			_, _ = conn.Write([]byte(msg)) // Retransmit
			continue
		}

		ack := string(buf[:n])
		fmt.Println("Received:", ack)
	}
}

func receiver3_0() {
	addr, _ := net.ResolveUDPAddr("udp", receiverAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	expectedSeq := 0

	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)
		message := string(buffer[:n])
		parts := strings.SplitN(message, "|", 2)

		// Simulate packet loss (10% chance)
		if rand.Intn(10) < 1 {
			fmt.Println("Simulated packet loss! Ignoring message.")
			continue
		}

		if len(parts) != 2 {
			fmt.Println("Corrupted packet! Resending last ACK.")
			ackMsg := fmt.Sprintf("ACK%d", expectedSeq^1)
			conn.WriteToUDP([]byte(ackMsg), remoteAddr)
			continue
		}

		seqNum := parts[0]
		data := parts[1]

		// Simulate corruption (10% chance)
		if rand.Intn(10) < 1 {
			fmt.Println("Simulated corruption! Resending last ACK.")
			ackMsg := fmt.Sprintf("ACK%d", expectedSeq^1)
			conn.WriteToUDP([]byte(ackMsg), remoteAddr)
			continue
		}

		// Check if expected sequence number matches
		if seqNum == fmt.Sprintf("%d", expectedSeq) {
			fmt.Println("Received:", data, "with seq:", seqNum)
			expectedSeq = 1 - expectedSeq // Toggle expected sequence
		} else {
			fmt.Println("Duplicate packet detected. Resending last ACK.")
		}

		// Send ACK for the last correctly received packet
		ackMsg := fmt.Sprintf("ACK%d", expectedSeq^1)
		conn.WriteToUDP([]byte(ackMsg), remoteAddr)
	}
}

func receiver2_1() {
	addr, _ := net.ResolveUDPAddr("udp", receiverAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	expectedSeq := 0

	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)
		message := string(buffer[:n])
		parts := strings.SplitN(message, "|", 2)

		if len(parts) != 2 {
			fmt.Println("Corrupted packet! Ignoring.")
			continue
		}

		seqNum := parts[0]
		data := parts[1]

		// Simulate garbled ACK (10% chance)
		if rand.Intn(10) < 1 {
			fmt.Println("Garbled ACK response!")
			conn.WriteToUDP([]byte("GARBLED"), remoteAddr)
			continue
		}

		// Check if expected sequence number matches
		if seqNum == fmt.Sprintf("%d", expectedSeq) {
			fmt.Println("Received:", data, "with seq:", seqNum)
			expectedSeq = 1 - expectedSeq // Toggle expected sequence
		} else {
			fmt.Println("Duplicate packet detected. Resending last ACK.")
		}

		// Send ACK for the last correctly received packet
		ackMsg := fmt.Sprintf("ACK%d", expectedSeq^1)
		conn.WriteToUDP([]byte(ackMsg), remoteAddr)
	}
}

func receiver2_0() {
	addr, _ := net.ResolveUDPAddr("udp", receiverAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	go func() {
		for {
			n, remoteAddr, _ := conn.ReadFromUDP(buffer)
			message := string(buffer[:n])
			fmt.Println("Received:", message)

			// Simulate errors (10% chance of corruption)
			if rand.Intn(10) < 1 {
				fmt.Println("Corrupted packet! Sending NAK")
				conn.WriteToUDP([]byte("NAK"), remoteAddr)
				continue
			}

			fmt.Println("Packet OK! Sending ACK")
			conn.WriteToUDP([]byte("ACK"), remoteAddr)
		}
	}()

}

const (
	maxPackets = 16   // Maximum packets to send
	rtt        = 100  // Simulated round-trip time in ms
	lossRate   = 0.15 // Simulated packet loss rate
)

// additive increase : increase sending rate by 1 maximum segement size every RRT until loss detected
// multiplicative decrease : cut sendint rate in half at each loss event
func tcpCongestionControlSender() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// TCP Congestion Control Variables
	cwnd := 1     // Congestion Window (starts at 1 segment)
	ssthresh := 8 // Slow Start Threshold
	dupACKCount := 0

	for sentPackets := 0; sentPackets < maxPackets; {
		fmt.Printf("\n[CWND: %d, SSTHRESH: %d]\n", cwnd, ssthresh)

		// Send packets in the congestion window
		for i := 0; i < cwnd && sentPackets < maxPackets; i++ {
			packet := fmt.Sprintf("Packet %d", sentPackets+1)
			_, err := conn.Write([]byte(packet))
			if err != nil {
				fmt.Println("Error sending data:", err)
				return
			}
			fmt.Println("Sent:", packet)
			sentPackets++
		}

		// Simulate RTT delay
		time.Sleep(rtt * time.Millisecond)

		// Receive ACKs
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading ACK:", err)
			return
		}

		ackMsg := string(buffer[:n])
		if ackMsg == "ACK" {
			fmt.Println("Received ACK")

			// Simulate packet loss scenario
			if math.Mod(float64(sentPackets), 4) == 0 {
				fmt.Println("Simulating packet loss...")
				dupACKCount++
				if dupACKCount >= 3 {
					// Fast Retransmit & Recovery
					fmt.Println("3 Duplicate ACKs detected, Fast Retransmit & Reduce CWND")
					ssthresh = max(1, cwnd/2) // Halve threshold
					cwnd = 1                  // Restart Slow Start
					dupACKCount = 0
				}
			} else {
				dupACKCount = 0
				// Congestion Control Logic
				if cwnd < ssthresh {
					cwnd *= 2 // Slow Start (Exponential Growth)
				} else {
					cwnd++ // Congestion Avoidance (Linear Growth)
				}
			}
		}
	}
	fmt.Println("Finished sending all packets.")
}

// Helper function to get max value
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tcpCongestionControlReceiver() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed by client")
			return
		}

		data := string(buffer[:n])
		fmt.Println("Received:", data)

		// Simulate ACK response
		_, err = conn.Write([]byte("ACK"))
		if err != nil {
			fmt.Println("Error sending ACK:", err)
			return
		}
	}
}

// ComputeChecksum calculates the TCP checksum
func ComputeChecksum(data []byte) uint16 {
	var sum uint32

	// Sum 16-bit words
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(binary.BigEndian.Uint16(data[i : i+2]))
	}

	// If odd length, add the last byte
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}

	// Fold sum to 16 bits
	for (sum >> 16) > 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}

	// One's complement
	return ^uint16(sum)
}

// CreatePseudoHeader constructs a TCP pseudo-header
func CreatePseudoHeader(srcIP, dstIP string, tcpLen uint16) []byte {
	pseudoHeader := make([]byte, 12)
	src := net.ParseIP(srcIP).To4()
	dst := net.ParseIP(dstIP).To4()

	copy(pseudoHeader[0:4], src)
	copy(pseudoHeader[4:8], dst)

	pseudoHeader[8] = 0                                   // Reserved
	pseudoHeader[9] = 6                                   // TCP Protocol number (6)
	binary.BigEndian.PutUint16(pseudoHeader[10:], tcpLen) // TCP Length

	return pseudoHeader
}

// ComputeTCPChecksum calculates the TCP checksum with the pseudo-header
func ComputeTCPChecksum(srcIP, dstIP string, tcpHeader, payload []byte) uint16 {
	tcpLen := uint16(len(tcpHeader) + len(payload))
	pseudoHeader := CreatePseudoHeader(srcIP, dstIP, tcpLen)

	// Concatenate pseudo-header, TCP header, and payload
	fullSegment := append(pseudoHeader, tcpHeader...)
	fullSegment = append(fullSegment, payload...)

	// Compute checksum
	return ComputeChecksum(fullSegment)
}

func checksumMain() {
	srcIP := "192.168.1.1"
	dstIP := "192.168.1.2"

	// Sample TCP header (20 bytes with some dummy values)
	tcpHeader := []byte{
		0x04, 0xD2, // Source Port: 1234
		0x00, 0x50, // Dest Port: 80
		0x12, 0x34, 0x56, 0x78, // Sequence Number
		0x9A, 0xBC, 0xDE, 0xF0, // Acknowledgment Number
		0x50, 0x02, // Header Length & Flags (SYN)
		0x72, 0x10, // Window Size
		0x00, 0x00, // Checksum (to be computed)
		0x00, 0x00, // Urgent Pointer
	}

	payload := []byte("Hello, TCP!") // Sample payload

	// Compute TCP checksum
	checksum := ComputeTCPChecksum(srcIP, dstIP, tcpHeader, payload)

	// Display result
	fmt.Printf("Computed TCP Checksum: 0x%X\n", checksum)

	// Insert the checksum into the TCP header
	binary.BigEndian.PutUint16(tcpHeader[16:], checksum)

	// Print final TCP segment with checksum
	fmt.Println("Final TCP Header with Checksum:", tcpHeader)
}

// Finite State Machine (FSM)

// WAIT_FOR_CALL_0, WAIT_FOR_ACK_0: Sender waiting to send packet 0 / ACK for 0.
// WAIT_FOR_CALL_1, WAIT_FOR_ACK_1: Sender waiting to send packet 1 / ACK for 1.
// Sender Function

// Sends packet with sequence number.
// Waits for ACK (handles timeout for retransmission).
// Uses FSM to track state transitions.
// Receiver Function

// Reads incoming packets.
// Checks for corruption (simulated).
// Sends ACK if packet is valid.
// Timeout Handling

// If no ACK is received within 2 seconds, packet is resent.
// Constants for FSM states
const (
	WAIT_FOR_CALL_0 = iota
	WAIT_FOR_ACK_0
	WAIT_FOR_CALL_1
	WAIT_FOR_ACK_1
)

// Packet structure
type Packet struct {
	SeqNum  int
	AckNum  int
	Payload string
}

// Simulate possible corruption
func corruptPacket() bool {
	return rand.Float32() < 0.1 // 10% chance of corruption
}

// Sender function (RDT 2.2 with FSM)
func sender(conn *net.UDPConn, addr *net.UDPAddr) {
	state := WAIT_FOR_CALL_0
	seqNum := 0

	for i := 0; i < 5; i++ { // Sending 5 packets
		packet := Packet{SeqNum: seqNum, Payload: fmt.Sprintf("Message %d", i)}
		data := fmt.Sprintf("%d:%s", packet.SeqNum, packet.Payload)

		// Send packet
		fmt.Printf("[SENDER] Sending: %s\n", data)
		_, err := conn.WriteToUDP([]byte(data), addr)
		if err != nil {
			fmt.Println("Error sending packet:", err)
			return
		}

		// Wait for ACK
		buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(2 * time.Second)) // Timer for ACK
		_, _, err = conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("[SENDER] Timeout! Resending...")
			i-- // Resend same packet
			continue
		}

		// Process ACK
		ackNum := int(buffer[0] - '0')
		if ackNum == seqNum {
			fmt.Printf("[SENDER] Received ACK: %d\n", ackNum)
			seqNum = 1 - seqNum // Flip sequence number
			state = (state + 1) % 4
		} else {
			fmt.Println("[SENDER] Wrong ACK! Resending...")
			i-- // Resend same packet
		}
	}
}

// Receiver function (RDT 2.2 FSM)
func receiver(conn *net.UDPConn) {
	state := WAIT_FOR_CALL_0
	expectedSeq := 0

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error receiving packet:", err)
			return
		}

		// Simulate corruption
		if corruptPacket() {
			fmt.Println("[RECEIVER] Corrupted packet received, ignoring.")
			continue
		}

		// Extract sequence number
		seqNum := int(buffer[0] - '0')
		message := string(buffer[2:n])

		// Process valid packet
		if seqNum == expectedSeq {
			fmt.Printf("[RECEIVER] Received: %s\n", message)
			ack := fmt.Sprintf("%d", expectedSeq)
			conn.WriteToUDP([]byte(ack), addr)
			expectedSeq = 1 - expectedSeq // Flip expected sequence
			state = (state + 1) % 4
		} else {
			fmt.Println("[RECEIVER] Duplicate packet received, resending last ACK.")
			ack := fmt.Sprintf("%d", 1-expectedSeq) // Send last valid ACK
			conn.WriteToUDP([]byte(ack), addr)
		}
	}
}

func FiniteStateMachinemain() {
	// UDP connection setup
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	go receiver(conn) // Start receiver

	time.Sleep(1 * time.Second) // Wait before starting sender

	// Connect sender to receiver
	senderAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	senderConn, _ := net.DialUDP("udp", nil, senderAddr)
	defer senderConn.Close()

	sender(senderConn, senderAddr)
}

// Fragmentation: Splitting large messages into fixed-size records.
// Compression (Optional): Not commonly used, but can be simulated.
// Encryption: Encrypt data using AES-GCM for confidentiality.
// MAC (Message Authentication Code): Ensures integrity and authenticity.
// Transmission & Decryption: Reverse process at the receiver.

// TLSRecord struct simulating a TLS record
type TLSRecord struct {
	Type    byte   // 23 for Application Data
	Version uint16 // TLS 1.2 = 0x0303
	Length  uint16
	Payload []byte
	Mac     []byte
}

// Generate random AES key
func generateAESKey() []byte {
	key := make([]byte, 32) // AES-256
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

// Generate random nonce for AES-GCM
func generateNonce() []byte {
	nonce := make([]byte, 12) // 96-bit nonce for AES-GCM
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	return nonce
}

// Encrypt data using AES-GCM
func encrypt(data, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesGCM.Seal(nil, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt data using AES-GCM
func decrypt(ciphertext, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Compute HMAC for message authentication
func computeHMAC(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// Simulate sending a TLS record
func sendTLSRecord(message string, key []byte) TLSRecord {
	nonce := generateNonce()
	encryptedPayload, _ := encrypt([]byte(message), key, nonce)
	mac := computeHMAC(encryptedPayload, key)

	record := TLSRecord{
		Type:    23,     // Application Data
		Version: 0x0303, // TLS 1.2
		Length:  uint16(len(encryptedPayload)),
		Payload: encryptedPayload,
		Mac:     mac,
	}

	fmt.Println("\n[SENDER] TLS Record Sent")
	fmt.Printf("Type: %d | Version: 0x%X | Length: %d\n", record.Type, record.Version, record.Length)
	fmt.Printf("Encrypted Payload: %s\n", hex.EncodeToString(record.Payload))
	fmt.Printf("MAC: %s\n", hex.EncodeToString(record.Mac))
	return record
}

// Simulate receiving and processing a TLS record
func receiveTLSRecord(record TLSRecord, key []byte) {
	fmt.Println("\n[RECEIVER] Processing TLS Record...")

	// Validate MAC
	expectedMac := computeHMAC(record.Payload, key)
	if !hmac.Equal(record.Mac, expectedMac) {
		fmt.Println("[ERROR] MAC verification failed! Possible data tampering detected.")
		return
	}

	// Decrypt payload
	nonce := generateNonce()
	decryptedPayload, err := decrypt(record.Payload, key, nonce)
	if err != nil {
		fmt.Println("[ERROR] Decryption failed!", err)
		return
	}

	fmt.Printf("[RECEIVER] Decrypted Message: %s\n", string(decryptedPayload))
}

func TLSRecordMain() {
	// Simulate TLS communication
	key := generateAESKey()
	message := "Hello, this is a secure TLS message."

	// Sender: Encrypt & Send
	tlsRecord := sendTLSRecord(message, key)

	// Receiver: Decrypt & Validate
	receiveTLSRecord(tlsRecord, key)
}

// To implement TLS handshake message types and protocol actions, you'll need to understand the basic flow of the TLS handshake process and the structure of each message type. Below is an outline of the key components involved in the implementation of TLS handshake message types and the corresponding actions that occur during the handshake:

// TLS Handshake Flow
// ClientHello: The client initiates the handshake by sending the ClientHello message, which includes details like supported cipher suites, TLS version, and random data.
// ServerHello: The server responds with a ServerHello message, which includes the selected cipher suite, TLS version, and random data.
// Server Certificate: The server sends its certificate to prove its identity.
// Key Exchange: Based on the cipher suite selected, the client and server exchange keying material to generate a shared secret (e.g., Diffie-Hellman key exchange).
// Server Finished: The server sends a Finished message to indicate it has completed the handshake process on its side.
// Client Finished: The client sends a Finished message to confirm it has completed the handshake.

// ClientHello Message: The client sends a message that starts with a byte representing the message type (ClientHello), followed by random data (which would be used in real handshakes for key generation).
// ServerHello Message: The server responds with its message starting with ServerHello, followed by random data to match the cipher suite and other parameters.
// Server Certificate: The server sends a certificate to prove its identity (in this simplified example, we are using random data as a placeholder).
// Key Exchange: Both the client and server exchange keys to generate a shared secret.
// Finished Messages: The client and server exchange Finished messages indicating that the handshake has been completed.

const (
	// Handshake message types
	ClientHello       = 1
	ServerHello       = 2
	ServerCertificate = 11
	ClientKeyExchange = 16
	ServerFinished    = 20
	ClientFinished    = 22
)

// Simple TLS Handshake
func tlsHandshake(conn net.Conn, isClient bool) {
	// Generate random data for the handshake
	randData := make([]byte, 32)
	_, err := rand.Read(randData)
	if err != nil {
		log.Fatalf("Failed to generate random data: %v", err)
	}

	if isClient {
		// ClientHello message (simplified version)
		clientHello := append([]byte{ClientHello}, randData...)
		_, err := conn.Write(clientHello)
		if err != nil {
			log.Fatalf("Failed to send ClientHello: %v", err)
		}
		fmt.Printf("ClientHello sent: %v\n", clientHello)

		// Wait for ServerHello
		serverHello := make([]byte, 1024)
		_, err = conn.Read(serverHello)
		if err != nil {
			log.Fatalf("Failed to read ServerHello: %v", err)
		}
		if serverHello[0] == ServerHello {
			fmt.Printf("ServerHello received: %v\n", serverHello)
		}

		// Receive Server Certificate
		serverCert := make([]byte, 1024)
		_, err = conn.Read(serverCert)
		if err != nil {
			log.Fatalf("Failed to read Server Certificate: %v", err)
		}
		fmt.Printf("Server Certificate received: %v\n", serverCert)

		// Simulate Client Key Exchange
		clientKeyExchange := append([]byte{ClientKeyExchange}, randData...)
		_, err = conn.Write(clientKeyExchange)
		if err != nil {
			log.Fatalf("Failed to send Client Key Exchange: %v", err)
		}
		fmt.Printf("Client Key Exchange sent: %v\n", clientKeyExchange)

		// Wait for Server Finished message
		serverFinished := make([]byte, 1024)
		_, err = conn.Read(serverFinished)
		if err != nil {
			log.Fatalf("Failed to read Server Finished: %v", err)
		}
		fmt.Printf("Server Finished received: %v\n", serverFinished)

		// Send Client Finished message
		clientFinished := []byte{ClientFinished}
		_, err = conn.Write(clientFinished)
		if err != nil {
			log.Fatalf("Failed to send Client Finished: %v", err)
		}
		fmt.Printf("Client Finished sent: %v\n", clientFinished)

	} else {
		// Server side of the handshake

		// Wait for ClientHello
		clientHello := make([]byte, 1024)
		_, err := conn.Read(clientHello)
		if err != nil {
			log.Fatalf("Failed to read ClientHello: %v", err)
		}
		if clientHello[0] == ClientHello {
			fmt.Printf("ClientHello received: %v\n", clientHello)
		}

		// Send ServerHello message
		serverHello := append([]byte{ServerHello}, randData...)
		_, err = conn.Write(serverHello)
		if err != nil {
			log.Fatalf("Failed to send ServerHello: %v", err)
		}
		fmt.Printf("ServerHello sent: %v\n", serverHello)

		// Send Server Certificate message (simplified)
		serverCertificate := append([]byte{ServerCertificate}, randData...)
		_, err = conn.Write(serverCertificate)
		if err != nil {
			log.Fatalf("Failed to send Server Certificate: %v", err)
		}
		fmt.Printf("Server Certificate sent: %v\n", serverCertificate)

		// Receive Client Key Exchange
		clientKeyExchange := make([]byte, 1024)
		_, err = conn.Read(clientKeyExchange)
		if err != nil {
			log.Fatalf("Failed to read Client Key Exchange: %v", err)
		}
		fmt.Printf("Client Key Exchange received: %v\n", clientKeyExchange)

		// Send Server Finished message
		serverFinished := []byte{ServerFinished}
		_, err = conn.Write(serverFinished)
		if err != nil {
			log.Fatalf("Failed to send Server Finished: %v", err)
		}
		fmt.Printf("Server Finished sent: %v\n", serverFinished)

		// Wait for Client Finished message
		clientFinished := make([]byte, 1024)
		_, err = conn.Read(clientFinished)
		if err != nil {
			log.Fatalf("Failed to read Client Finished: %v", err)
		}
		fmt.Printf("Client Finished received: %v\n", clientFinished)
	}
}

func main() {
	// Server setup
	serverAddr := "localhost:4433"
	server, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Server started. Waiting for client...")

	// Accepting client connection
	conn, err := server.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()

	// Perform TLS handshake as server
	//tlsHandshake(conn, isClient: false)
}
