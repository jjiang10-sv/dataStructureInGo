package numerical

import "testing"

// Testfactorial tests factorial0521 function
func Test_factorial(t *testing.T) {

	if factorial0521(2) != 2 {
		t.Error("[Error] factorial0521(2) is wrong")
	}

	if factorial0521(3) != 6 {
		t.Error("[Error] factorial0521(3) is wrong")
	}

	if factorial0521(0) != 1 {
		t.Error("[Error] factorial0521(0) is wrong")
	}

	if factorial0521(5) != 120 {
		t.Error("[Error] factorial0521(5) is wrong")
	}
}
