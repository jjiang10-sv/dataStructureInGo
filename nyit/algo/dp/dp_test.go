package main

import (
	"testing"
)

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_MultiStageFlow(t *testing.T) {
	mainFlow()
}
