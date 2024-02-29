package main

import (
	"time"
)

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	for num := 2; num < N; num++ {
		select {
		case <-stop:
			break
		default:
			if isPrime(num) {
				prime_nums <- num
				time.Sleep(time.Millisecond)
			}
		}
	}
	close(prime_nums)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
