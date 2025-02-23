package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// GenerateKeyPair generates a new ECDSA private and public key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// SignMessage signs a message using the private key
func SignMessage(privateKey *ecdsa.PrivateKey, message string) ([]byte, []byte, error) {
	hash := sha256.Sum256([]byte(message))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, nil, err
	}
	return r.Bytes(), s.Bytes(), nil
}

// VerifySignature verifies the signature using the public key
func VerifySignature(publicKey *ecdsa.PublicKey, message string, rBytes, sBytes []byte) bool {
	hash := sha256.Sum256([]byte(message))

	var r, s big.Int
	r.SetBytes(rBytes)
	s.SetBytes(sBytes)

	return ecdsa.Verify(publicKey, hash[:], &r, &s)
}

func mainSign() {
	// Generate key pair
	privateKey, err := GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Sign a message
	message := "Hello, secure world!"
	r, s, err := SignMessage(privateKey, message)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}

	fmt.Println("Signature (r, s):", r, s)

	// Verify the signature
	valid := VerifySignature(&privateKey.PublicKey, message, r, s)
	if valid {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature verification failed.")
	}
}
