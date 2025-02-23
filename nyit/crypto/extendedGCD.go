package main

import (
	"fmt"
)

// ExtendedEuclidean computes the GCD and finds x, y such that ax + by = gcd(a, b)
func ExtendedEuclidean(a, b int) (int, int, int, bool) {
	if b == 0 {
		return a, 1, 0, true
	}

	gcd, x1, y1, _ := ExtendedEuclidean(b, a%b)
	if gcd != 1 {
		return 0, 0, 0, false
	}
	x := y1
	y := x1 - (a/b)*y1

	return gcd, x, y, true

}

// ModInverse finds the modular inverse of a mod m using Extended Euclidean Algorithm
func ModInverse(a, m int) (int, error) {
	_, x, _, found := ExtendedEuclidean(a, m)

	// Inverse exists only if gcd(a, m) == 1
	if !found {
		return 0, fmt.Errorf("modular inverse does not exist for %d mod %d", a, m)
	}

	// x might be negative, so make it positive by adding m
	return (x%m + m) % m, nil
}
