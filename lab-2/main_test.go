package main

import (
	"testing"
	"time"
)

func Test_OrdersFinished(t *testing.T) {
	EAT_TIME = time.Second * 0
	THINK_TIME = time.Second * 0

	main()

	ordersFinished := len(Stats.philosophersIdx)
	totalOrders := len(Philosophers)

	if ordersFinished != totalOrders {
		t.Errorf("Not all orders finished. Expected %d got %d", totalOrders, ordersFinished)
	}

	Stats.philosophersIdx = []int{}

}

func Test_VarDelays(t *testing.T) {
	tests := []struct {
		name  string
		delay time.Duration
	}{
		{"NO_DELAY", time.Second * 0},
		{"0.25s_DELAY", time.Millisecond * 250},
		{"0.5s_DELAY", time.Millisecond * 500},
		{"0.75s_DELAY", time.Millisecond * 750},
	}

	for _, test := range tests {
		EAT_TIME = test.delay
		THINK_TIME = test.delay
		Test_OrdersFinished(t)
	}
}
