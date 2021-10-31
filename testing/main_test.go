package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 11 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
	// }

	tables := []struct {
		a int
		b int
		n int
	} {
		{ 1, 2, 3 },
		{ 2, 2, 4 },
		{ 25, 26, 51 },
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	} {
		{ 4, 2, 4 },
		{ 3, 2, 3 },
		{ 25, 26, 26 },
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if(max != item.n) {
			t.Errorf("Max was incorrect, got %d expected %d", max, item.n)
		}
	}
}

// go test
// go test -cover
// go test -coverprofile=cover.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a int
		n int
	} {
		{ 0, 0 },
		{ 1, 1 },
		{ 8, 21 },
		{ 50, 12586269025 },
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}

// go test -cpuprofile=cpu.out
// go tool pprof cpu.out
// (pprof) top
// (pprof) list Fibonacci
// (pprof) web
// (pprof) pdf