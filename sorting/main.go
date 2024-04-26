package main

import "fmt"

func sliceToChannel(nums []int) chan int {

	res := make(chan int)

	go func() {
		for _, i := range nums {
			res <- i
		}
		close(res)
	}()
	return res

}

func sq(in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for i := range in {
			res <- i * i
		}
		close(res)
	}()
	return res
}

func gotest() {

	data := []int{1, 2, 3, 4}

	sliceChan := sliceToChannel(data)
	sqChan := sq(sliceChan)

	for i := range sqChan {
		fmt.Println(i)
	}
}
