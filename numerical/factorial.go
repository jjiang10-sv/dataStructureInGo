package numerical

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func factorial0521(num int) int {
	
	if num == 0 {
		return 1
	}

	return num*factorial0521(num-1)
}