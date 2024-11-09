package main

/*
 * Merge sort - http://en.wikipedia.org/wiki/Merge_sort
 */

func Merge(left, right []int) []int {
    result := make([]int, 0, len(left) + len(right))
    
    for len(left) > 0 || len(right) > 0 {
        if len(left) == 0 {
            return append(result, right...)
        }
        if len(right) == 0 {
            return append(result, left...)
        }
        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    
    return result
}

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    middle := len(arr) / 2
    
    left := MergeSort(arr[:middle])
    right := MergeSort(arr[middle:])
    
    return Merge(left, right)
}

func mergetwoSortedArr(a,b []int) []int{
    
    lenA,lenB := len(a),len(b)
    // if lenA == 0 {
    //     return b
    // } else if lenB == 0 {
    //     return a
    // }
    l := lenA+lenB
    res := make([]int, l)
    for i:=0;i<l;i++{
        if len(a) == 0 {
            for j, item := range b {
                res[i+j] = item
            }
            break
        }else if len(b) == 0 {
            for j, item := range a {
                res[i+j] = item
            }
            break
        }
        tmpA,tmpB := a[0],b[0]
        if tmpA < tmpB {
            res[i] = tmpA
            a = a[1:]
            continue
        }
        res[i] = tmpB
        b = b[1:]
    }
    return res
}

func mergeSort0429(arr []int) []int{
    if len(arr) <= 1{
        return arr
    }

    mid := len(arr)/2
    left := mergeSort0429(arr[0:mid])
    right := mergeSort0429(arr[mid:])

    return mergeTwoSortedArrRecursive(left,right)
}

func mergeTwoSortedArrRecursive(a,b []int) []int{
    
    res := make([]int, 0, len(a)+len(b))
    recursiveGet(a,b,&res)
    return res

}
func recursiveGet(a,b []int, res *[]int) {
    if len(a) == 0 {
        *res = append(*res, b...)
    }else if len(b) == 0{
        *res = append(*res, a...)
    }else {
        min,minA,minB :=0, a[0],b[0]
        if minA < minB {
            a = a[1:]
            min = minA
        }else {
            min = minB
            b = b[1:]
        }
        *res = append(*res, min)
        recursiveGet(a,b,res)
    }
}