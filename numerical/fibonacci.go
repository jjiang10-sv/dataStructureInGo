package numerical

//using recursion O(2 power n)
func fibo(num int) int {
	if num <= 1 {
		return num
	}
	return fibo(num-1) + fibo(num-2)
}
// O(n)
func fiboIter0521(num int) int {
	if num <= 1 {
		return num
	}
	res, first,second :=0, 0,1
	for num > 1 {
		res = (first+second)
		first = second
		second = res
		num--
	}
	return res
}

func fiboDpGetSeq0521(num int) []int {
	if num < 0 {
		panic("argument shall be above 0")
	}
	res := make([]int, num+1)
	res[0] = 0
	if num >= 1 {
		res[1] = 1
	}
	if num <= 1 {
		return res
	}
	for i:=2; i <= num; i++{
		res[i] = (res[i-1]+res[i-2])
	}
	return res
}