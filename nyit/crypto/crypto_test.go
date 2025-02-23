package main

import (
	"fmt"
	"testing"
)

func TestDecryptSim(t *testing.T) {
	//decrypTable := getConstructionTable()
	decrypTable := getDecryptTableComp()
	// limited to uppercase
	ciphertext := "JHJKHasad"
	plaintext := decrypt(ciphertext, decrypTable)
	fmt.Println(plaintext)
}

func TestModInverse(t *testing.T) {
	// 4321, 1234
	divident, modular := 5, 3
	//a, m := modular, divident // Example: Find 3⁻¹ mod 7
	inverse, err := ModInverse(modular, divident)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Modular inverse of %d mod %d is %d\n", divident, modular, inverse)
	}
}

func TestMillerRabin(t *testing.T) {
	mainMiller()
}
