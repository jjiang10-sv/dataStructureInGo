package numerical

import (
	"fmt"
	"testing"
)

// TestGcd tests gcd
func TestGcd(t *testing.T) {

	// if GCD0521_2(100, 200) != 100 {
	// 	t.Error("[Error] GCD0521_2(100, 200) is wrong")
	// }

	// if GCD0521_2(200, 100) != 100 {
	// 	t.Error("[Error] GCD0521_2(100, 200) is wrong")
	// }

	// if GCD0521_2(4, 2) != 2 {
	// 	t.Error("[Error] GCD0521_2(4,2) is wrong")
	// }

	// if GCD0521_2(6, 3) != 3 {
	// 	t.Error("[Error] GCD0521_2(6,3) is wrong")
	// }

	// if GCD0521_2(4,6) != 2 {
	// 	t.Error(("wrong"))
	// }

	array := []int{-3, -4, 7, 1, -2, 0, -5, 1, 0, 6, -5}
	fmt.Println("Maximum subarray sum: ", maxSubarray(array))

	// array = []int{3, 4, -7, 2, 0, 0, -3, -1, 0, -5, 7}
	// fmt.Println("Maximum subarray sum: ", maxSubarray(array))
}
