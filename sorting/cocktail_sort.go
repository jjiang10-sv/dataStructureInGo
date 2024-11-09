package main

/*
 * Cocktail sort - https://en.wikipedia.org/wiki/Cocktail_sort
 */

func CocktailSort(arr []int) {
	tmp := 0

	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1

		for left <= right {

			if arr[left] > arr[left+1] {
				tmp = arr[left]
				arr[left] = arr[left+1]
				arr[left+1] = tmp
			}
			left++
			if arr[right-1] > arr[right] {
				tmp = arr[right-1]
				arr[right-1] = arr[right]
				arr[right] = tmp
			}
			right--
		}
	}
}

func CocktailSort0428(arr []int) {

	length := len(arr)

	for i:=0;i<length/2;i++{
		left:=0
		right:=length-1
		for left < right {
			if arr[left] > arr[left+1]{
				arr[left],arr[left+1] = arr[left+1],arr[left]
			}
			left++
			if arr[right]<arr[right-1]{
				arr[right],arr[right-1] = arr[right-1],arr[right]
			}
			right--
		}
	}
}