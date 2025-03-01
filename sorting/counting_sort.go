package main

/*
 * Counting sort - https://en.wikipedia.org/wiki/Counting_sort
 */

func getK(arr []int) int {
    if len(arr) == 0 {
        return 1
    }

    k := arr[0]
    
    for _, v := range arr {
        if v > k {
            k = v
        }
    }
    
    return k+1
}
// memory cost increase on the maxVal of the arr
// use case : smaill value like less than 10
// time cost O(n)
func CountingSort(arr []int) {
    k := getK(arr)
    array_of_counts := make([]int, k)

    for i:= 0; i < len(arr); i++ {
        array_of_counts[arr[i]] += 1
    }

    for i, j := 0, 0; i < k; i++ {
        for {
            if array_of_counts[i] > 0 {
                arr[j] = i
                j += 1
                array_of_counts[i] -= 1
                continue
            }
            break
        }
    }
}

func CountingSort0428(arr []int, max int) {
    countsArr := make([]int,max)
    for _, v := range arr{
        countsArr[v] +=1
    }
    for i,j:=0,0; i<max;i++{
        for {
            if countsArr[i] > 0{
                arr[j] = i
                j++
                countsArr[i]--
                continue
            }
            break
        }
        
    }
}