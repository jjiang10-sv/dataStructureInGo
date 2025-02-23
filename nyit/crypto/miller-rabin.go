package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// MillerRabinTest performs the Miller-Rabin primality test.
// n: The number to test for primality.
// k: The number of iterations (higher k increases accuracy).
// Returns true if n is probably prime, false if n is composite.
func MillerRabinTest(n, k int) bool {
	// Handle small numbers and even numbers
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Decompose n-1 into d * 2^s
	s := 0
	d := n - 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	// Perform k iterations
	for i := 0; i < k; i++ {
		// Pick a random integer a in the range [2, n-2]
		a := rand.Intn(n-3) + 2

		// Compute x = a^d mod n
		x := modularExponentiation(a, d, n)

		// Check if x is 1 or n-1
		if x == 1 || x == n-1 {
			continue
		}

		// Repeat squaring x up to s-1 times
		for j := 0; j < s-1; j++ {
			x = modularExponentiation(x, 2, n)
			if x == n-1 {
				break
			}
			if x == 1 {
				return false
			}
		}

		// If x is not n-1, n is composite
		if x != n-1 {
			return false
		}
	}

	// If no witness found, n is probably prime
	return true
}

// modularExponentiation computes (base^exponent) % modulus efficiently.
func modularExponentiation(base, exponent, modulus int) int {
	if modulus == 1 {
		return 0
	}
	result := 1
	base = base % modulus
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}

func modularExponentiationJohn(base, exponent, modular int) int {
	if modular == 1 {
		return 0
	}
	result := 1
	base = base % modular
	for exponent > 0 {
		// multiply by base if exponent is odd
		if exponent&1 == 1 {
			result = (result * base) % modular
		}
		exponent = exponent >> 1
		base = (base * base) % modular

	}
	return result
}

func millerRobinTestJohn(n, k int) bool {
	// n can not be less than 2 to be a prime
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	//if n is even, then not a prime
	if n&1 == 0 {
		return false
	}

	// if n%2 == 0 {
	// 	return false
	// }

	s, d := 0, n-1
	//run till d is odd
	for d&1 == 0 {
		d = d >> 1
		s++
	}

	// for d%2 == 0 {
	// 	d /= 2
	// 	s++
	// }

	for i := 0; i < k; i++ {
		// choose a random int in [2,n-2]
		a := rand.Intn(n-3) + 2
		//x := modularExponentiationJohn(a, d, n)
		x := modularExponentiation(a, d, n)
		if x == 1 || x == (n-1) {
			continue
		}
		// repeating square x up to s-1 time
		for s > 1 {
			s--
			x = modularExponentiation(a, 2, n)

			// if met, then go to next iteration in k
			if x == n-1 {
				break
			}
			// check x against 1; directly return false if met
			if x == 1 {
				return false
			}

		}
		// n is not a prime after iterate s -1 time
		if x != (n - 1) {
			return false
		}
	}
	return true

}
func mainMiller() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Test numbers for primality
	// 2, 3,
	numbers := []int{5, 7, 11, 13, 15, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 100, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

	for _, num := range numbers {
		isPrime := MillerRabinTest(num, 5) // 5 iterations for accuracy
		isPrimeJohn := millerRobinTestJohn(num, 5)
		print(isPrime == isPrimeJohn)
		fmt.Printf("%d is prime: %v\n", num, isPrime)
	}
}

// MillerRabinTest performs the Miller-Rabin primality test.
// n: The number to test for primality.
// k: The number of iterations (higher k increases accuracy).
// Returns true if n is probably prime, false if n is composite.
func MillerRabinTestBig(n *big.Int, k int) bool {
	// Handle small numbers and even numbers
	if n.Cmp(big.NewInt(2)) == -1 {
		return false
	}
	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	// Decompose n-1 into d * 2^s
	s := 0
	d := new(big.Int).Sub(n, big.NewInt(1))
	for new(big.Int).Mod(d, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		d.Div(d, big.NewInt(2))
		s++
	}

	// Perform k iterations
	for i := 0; i < k; i++ {
		// Pick a random integer a in the range [2, n-2]
		a := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(n, big.NewInt(3)))
		a.Add(a, big.NewInt(2))

		// Compute x = a^d mod n
		x := new(big.Int).Exp(a, d, n)

		// Check if x is 1 or n-1
		if x.Cmp(big.NewInt(1)) == 0 || x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
			continue
		}

		// Repeat squaring x up to s-1 times
		for j := 0; j < s-1; j++ {
			x.Exp(x, big.NewInt(2), n)
			if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) == 0 {
				break
			}
			if x.Cmp(big.NewInt(1)) == 0 {
				return false
			}
		}

		// If x is not n-1, n is composite
		if x.Cmp(new(big.Int).Sub(n, big.NewInt(1))) != 0 {
			return false
		}
	}

	// If no witness found, n is probably prime
	return true
}

func mainMillerBig() {
	// Test numbers for primality
	numbers := []int64{2, 3, 5, 7, 11, 13, 15, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 100, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}

	for _, num := range numbers {
		n := big.NewInt(num)
		isPrime := MillerRabinTestBig(n, 5) // 5 iterations for accuracy
		fmt.Printf("%d is prime: %v\n", num, isPrime)
	}
}

// DiscreteLogarithm computes the discrete logarithm of h with respect to g modulo p.
// Returns -1 if no solution exists.
func DiscreteLogarithm(g, h, p int) int {
	// Convert inputs to big.Int for modular arithmetic
	bigG := big.NewInt(int64(g))
	bigH := big.NewInt(int64(h))
	bigP := big.NewInt(int64(p))

	// Brute-force search for x
	x := 1
	for x < p {
		// Compute g^x mod p
		bigX := big.NewInt(int64(x))
		computedH := new(big.Int).Exp(bigG, bigX, bigP)

		// Check if computedH == h
		if computedH.Cmp(bigH) == 0 {
			return x
		}

		x++
	}

	// No solution found
	return -1
}

func maindisBig() {
	// Example: Find x such that 3^x ≡ 6 mod 7
	g := 3
	h := 6
	p := 7

	x := DiscreteLogarithm(g, h, p)
	if x != -1 {
		fmt.Printf("Discrete logarithm of %d with respect to %d modulo %d is: %d\n", h, g, p, x)
	} else {
		fmt.Println("No solution exists.")
	}
}

// DiscreteLogarithm computes the discrete logarithm of h with respect to g modulo p.
// Returns -1 if no solution exists.
func DiscreteLogarithmSmall(g, h, p int) int {
	// Initialize result to 1 (g^0 ≡ 1 mod p)
	result := 1

	// Iterate through all possible exponents x from 1 to p-1
	for x := 1; x < p; x++ {
		// Compute g^x mod p
		result = (result * g) % p

		// Check if g^x ≡ h mod p
		if result == h {
			return x
		}
	}

	// No solution found
	return -1
}

func mainDisSmall() {
	// Example: Find x such that 3^x ≡ 6 mod 7
	g := 3
	h := 6
	p := 7

	x := DiscreteLogarithmSmall(g, h, p)
	if x != -1 {
		fmt.Printf("Discrete logarithm of %d with respect to %d modulo %d is: %d\n", h, g, p, x)
	} else {
		fmt.Println("No solution exists.")
	}
}

// ModularExponentiation computes (base^exponent) % modulus efficiently.
func ModularExponentiation(base, exponent, modulus int) int {
	result := 1
	base = base % modulus
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}

// ModularInverse computes the modular inverse of a modulo m using Fermat's Little Theorem.
// This works only if m is prime.
func ModularInverse(a, m int) int {
	return ModularExponentiation(a, m-2, m)
}

// BabyStepGiantStep computes the discrete logarithm of h with respect to g modulo p.
// Returns -1 if no solution exists.
func BabyStepGiantStep(g, h, p int) int {
	// Compute m = ceil(sqrt(p))
	m := int(math.Ceil(math.Sqrt(float64(p))))

	// Baby Steps: Precompute g^j mod p for all 0 <= j < m
	babySteps := make(map[int]int)
	current := 1
	for j := 0; j < m; j++ {
		babySteps[current] = j
		current = (current * g) % p
	}

	// Compute g^{-m} mod p
	gInverseM := ModularExponentiation(ModularInverse(g, p), m, p)

	// Giant Steps: Compute h * (g^{-m})^i mod p for all 0 <= i < m
	current = h
	for i := 0; i < m; i++ {
		// Check if current exists in the babySteps map
		if j, exists := babySteps[current]; exists {
			// Solution found: x = i * m + j
			return i*m + j
		}
		// Update current for the next iteration
		current = (current * gInverseM) % p
	}

	// No solution found
	return -1
}

func mainBabyBigStep() {
	// Example: Find x such that 3^x ≡ 6 mod 7
	g := 3
	h := 6
	p := 7

	x := BabyStepGiantStep(g, h, p)
	if x != -1 {
		fmt.Printf("Discrete logarithm of %d with respect to %d modulo %d is: %d\n", h, g, p, x)
	} else {
		fmt.Println("No solution exists.")
	}
}
