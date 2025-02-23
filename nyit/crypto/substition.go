package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	// Step 4: Perform the chosen plaintext attack
// 	reconstructedTable := chosenPlaintextAttack()

// 	// Print the reconstructed substitution table
// 	fmt.Println("Reconstructed Substitution Table:")
// 	for c := 'A'; c <= 'Z'; c++ {
// 		fmt.Printf("%c -> %c\n", c, reconstructedTable[c])
// 	}

// 	// Step 5: Decrypt an intercepted ciphertext
// 	interceptedCiphertext := "QWERTYUIOPASDFGHJKLZXCVBNM"
// 	decryptedText := decrypt(interceptedCiphertext, reconstructedTable)
// 	fmt.Printf("\nDecrypted Text: %s\n", decryptedText)
// }

func encrypt(plaintext string) string {

	// Simulated substitution table (unknown to the attacker)
	var substitutionTable = map[rune]rune{
		'A': 'Q', 'B': 'W', 'C': 'E', 'D': 'R', 'E': 'T',
		'F': 'Y', 'G': 'U', 'H': 'I', 'I': 'O', 'J': 'P',
		'K': 'A', 'L': 'S', 'M': 'D', 'N': 'F', 'O': 'G',
		'P': 'H', 'Q': 'J', 'R': 'K', 'S': 'L', 'T': 'Z',
		'U': 'X', 'V': 'C', 'W': 'V', 'X': 'B', 'Y': 'N',
		'Z': 'M',
	}
	var res strings.Builder
	for _, char := range plaintext {
		res.WriteRune(substitutionTable[char])
	}
	return res.String()
}

func getConstructionTable() map[rune]rune {
	res := make(map[rune]rune)
	// attack submits the entire alphabets to the encrypt scheme
	plaintext := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ciphertext := encrypt(plaintext)
	for i := 0; i < len(plaintext); i++ {
		res[rune(ciphertext[i])] = rune(plaintext[i])
	}
	return res
}

func decrypt(cipherText string, decryptTable map[rune]rune) string {
	var plaintextBuilder strings.Builder
	for _, c := range cipherText {
		fmt.Println(string(c), string(decryptTable[c]))
		if plainRune, exist := decryptTable[c]; exist {
			plaintextBuilder.WriteRune(plainRune)
		} else {
			plaintextBuilder.WriteRune(c)
		}

	}
	return plaintextBuilder.String()
}

func getDecryptTableComp() map[rune]rune {
	constructTable, decrypTable := make(map[rune]rune), make(map[rune]rune)

	// construct from A to Z and repeated construct the guessed substitution table
	for c := 'A'; c <= 'Z'; c++ {
		plaintext := strings.Repeat(string(c), 10)
		ciphertext := encrypt(plaintext)
		for i, char := range ciphertext {
			constructTable[c+rune(i)] = char
		}
	}
	// reverse it to get decryptTable
	for k, v := range constructTable {
		decrypTable[v] = k
	}
	return decrypTable
}
