package main

/*
 * Quick sort - https://en.wikipedia.org/wiki/Quicksort
 */

import "math/rand"

func quick_sort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = quick_sort(low_part)
	high_part = quick_sort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}
func quickSort0430(arr []int) []int{
	if len(arr) <=1 {
		return arr
	}
	base := arr[rand.Intn(len(arr))]
	lowpart,midPart,highPart := []int{},[]int{},[]int{}
	for item,_ := range arr {
		if item > base{
			lowpart = append(lowpart, item)
		} else if item == base {
			midPart = append(midPart, item)
		}else {
			highPart = append(highPart, item)
		}
	}
	lowpart = quickSort0430(lowpart)
	highPart = quickSort0430(highPart)
	lowpart = append(lowpart, midPart...)
	lowpart = append(lowpart, highPart...)
	return lowpart

}