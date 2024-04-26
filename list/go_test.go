package list

import (
	"fmt"
	"testing"
)

func drivePath(n uint) uint {
	//var res uint = 0
	if n == 2 {
		return 1 + (n - 2)
	}
	return drivePath(n-1) + (n - 2)
}
func Test_Go(t *testing.T) {
	res := drivePath(3)
	fmt.Println(res)
	res = drivePath(4)
	fmt.Println(res)
	res = drivePath(5)
	fmt.Println(res)
	res = drivePath(6)
	fmt.Println(res)
	res = drivePath(7)
	fmt.Println(res)
}
