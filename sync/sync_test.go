package thres_sync

import (
	"testing"
)

func compare(x interface{}, y interface{}) bool {
	return x.(int) < y.(int)
}

func Test_diningPhy(t *testing.T) {
	//dining()
	produceAndConsume()
}
