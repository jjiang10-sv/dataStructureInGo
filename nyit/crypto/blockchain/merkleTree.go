package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// HashData hashes input string using SHA-256
func HashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// MerkleTree struct
type MerkleTree struct {
	Transactions []string
	Root         string
}

// NewMerkleTree constructs a Merkle Tree and returns the root hash
func NewMerkleTree(transactions []string) *MerkleTree {
	root := BuildMerkleTree(transactions)
	return &MerkleTree{Transactions: transactions, Root: root}
}

// BuildMerkleTree recursively computes the Merkle Root
func BuildMerkleTree(nodes []string) string {
	if len(nodes) == 1 {
		return nodes[0] // Root reached
	}

	// Ensure even number of nodes (duplicate last if odd)
	if len(nodes)%2 != 0 {
		nodes = append(nodes, nodes[len(nodes)-1])
	}

	var parentLevel []string
	for i := 0; i < len(nodes); i += 2 {
		concatenated := nodes[i] + nodes[i+1]
		parentLevel = append(parentLevel, HashData(concatenated))
	}

	return BuildMerkleTree(parentLevel)
}

// Main function
func mainMerkle() {
	transactions := []string{"tx1", "tx2", "tx3", "tx4", "tx5"}

	merkleTree := NewMerkleTree(transactions)

	fmt.Println("Merkle Root:", merkleTree.Root)
}

type Node struct {
	value string
}

type merkeNodeTree struct {
	transactions []Node
	root string
}

func hashNode(node Node) (string, error) {
	nodeData, err := json.Marshal(node)
	if err != nil {
		return "",err
	}
	nodeDataHash := sha256.Sum256(nodeData)
	return hex.EncodeToString(nodeDataHash[:]),nil
}

//func 