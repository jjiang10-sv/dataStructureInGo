package main

import "fmt"

// Function to generate a magic square of odd order n
func generateMagicSquare(n int) [][]int {
	// Initialize an empty 2D slice to hold the magic square
	magicSquare := make([][]int, n)
	for i := 0; i < n; i++ {
		magicSquare[i] = make([]int, n)
	}

	// Start placing 1 in the middle of the top row
	num := 1
	i, j := 0, n/2 // i is row, j is column

	for num <= n*n {
		magicSquare[i][j] = num
		num++

		// Move to the "next" position
		newI, newJ := (i-1+n)%n, (j+1)%n // Wrap around using modulus

		// If the new position is already occupied, move down instead
		if magicSquare[newI][newJ] != 0 {
			i++
		} else {
			i, j = newI, newJ
		}
	}

	return magicSquare
}

// Function to print the magic square
func printMagicSquare(magicSquare [][]int) {
	for i := 0; i < len(magicSquare); i++ {
		for j := 0; j < len(magicSquare[i]); j++ {
			fmt.Printf("%2d ", magicSquare[i][j])
		}
		fmt.Println()
	}
}

func mainMagic() {
	// Example: Generate a 5x5 magic square
	n := 5
	magicSquare := generateMagicSquare(n)
	printMagicSquare(magicSquare)
}