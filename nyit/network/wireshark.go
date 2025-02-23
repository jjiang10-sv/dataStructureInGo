package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device      = "eth0" // Change this to your network interface
	snapshotLen = int32(1600)
	promiscuous = false
	timeoutW    = 30 * time.Second
	handle      *pcap.Handle
	err         error
)

func wireshark() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeoutW)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter (optional)
	var filter string = "tcp and port 80" // Example: Capture only TCP traffic on port 80
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an Ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Printf("Source MAC: %s\n", ethernetPacket.SrcMAC)
		fmt.Printf("Destination MAC: %s\n", ethernetPacket.DstMAC)
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol:", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number:", tcp.Seq)
		fmt.Println()
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}
