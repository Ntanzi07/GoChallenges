package services

import (
	"Ex2/internal/data"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func CalculaFib(n int) (data.FibResult, bool) {
	if r, ok := data.GetResult(n); ok {
		return r, true
	}

	start := time.Now()
	resultChan := make(chan int)

	go func() {
		resultChan <- fib(n)
	}()

	select {
	case result := <-resultChan:
		duration := time.Since(start).String()
		fibResult := data.FibResult{
			Input:    n,
			Result:   result,
			Duration: duration,
		}
		data.SaveResult(fibResult)
		return fibResult, true
	case <-time.After(500 * time.Millisecond):
		return data.FibResult{}, false
	}
}
