package numerical

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	if fiboIter0521(10) != 55 {
		t.Error("[Error] Fibonacci(10) is wrong")
	}

	if fiboIter0521(0) != 0 {
		t.Error("[Error] Fibonacci(0) is wrong")
	}

	if fiboIter0521(3) != 2 {
		t.Error("[Error] Fibonacci(3) is wrong")
	}

	fmt.Println(fiboDpGetSeq0521(10))

}
