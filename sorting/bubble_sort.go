package main

/*
 * Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
 */

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

func BubbleSourt0428(arr []int) {
	length := len(arr)
	for i:=0; i<length;i++{
		for j:=0;j<length-i-1;j++{
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}