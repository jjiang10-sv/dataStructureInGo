package sortComplex

func partition(arr []int, low int, high int) int {

	pivot := arr[low]
	i := low + 1
	j := high
	for i < j {
		for i < j && arr[i] <= pivot {
			i++
		}
		for i < j && arr[j] >= pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[i] > pivot {
		i--
	}
	arr[low], arr[i] = arr[i], arr[low]
	return i
}


func partitionMid(arr []int, low int, high int) int {
	mid := (low + high) / 2
	pivot := arr[mid]
	i := low
	j := high
	for i < j {
		for i < high && arr[i] <= pivot {
			i++
		}
		for j > low && arr[j] >= pivot {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[i] > pivot {
		i--
	}
	arr[mid], arr[i] = arr[i], arr[mid]
	return i
}

func partitionMidSort(arr []int, low int, high int, isDesc bool) int {
	mid := (low + high) / 2
	pivot := arr[mid]
	i := low
	j := high
	for i < j {
		if isDesc {
			for i < j && arr[i] >= pivot {
				i++
			}
			for i < j && arr[j] <= pivot {
				j--
			}
		} else {
			for i < j && arr[i] <= pivot {
				i++
			}
			for i < j && arr[j] >= pivot {
				j--
			}
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// arr := []int{10, 7, 8, 9, 1, 5}
	if isDesc {
		if arr[i] < pivot {
			i--
		}
	} else {
		if arr[i] > pivot {
			i--
		}
	}
	arr[mid], arr[i] = arr[i], arr[mid]
	return i
}

func quick_sort(arr []int, low, high int) {

	if low < high {
		//pivot := partitionMidSort(arr, low, high, false)
		pivot := partitionMid(arr, low, high)
		
		quick_sort(arr, low, pivot-1)
		quick_sort(arr, pivot+1, high)
	}
}
