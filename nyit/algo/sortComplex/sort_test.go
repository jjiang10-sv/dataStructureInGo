package sortComplex

import (
	"fmt"
	"testing"
)

func Test_quick_sort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5,10, 7, 8, 9, 1, 5}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
