package main

import (
	"crypto/sha256"
	"fmt"
)

// Hash function using SHA-256
func hash(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func main() {
	data := "hello world"
	hashedValue := hash(data)
	fmt.Println("Hashed Value:", hashedValue)
}