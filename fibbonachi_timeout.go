package main

import (
	"errors"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	ch := make(chan int)

	go func() {
		ch <- fibonacci(n)
	}()

	select {
	case result := <-ch:
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout exceeded")
	}
}
