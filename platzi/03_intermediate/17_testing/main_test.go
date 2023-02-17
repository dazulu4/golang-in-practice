package main

import "testing"

// Individual Test
// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)
// 	if total != 10 {
// 		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
// 	}
// }

// Table Test Cases
func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d, expected %d", fib, item.n)
		}
	}
}

//go test -coverprofile=coverage.out
//go tool cover -func=coverage.out
//go tool cover -html=coverage.out

//profiling
//go test -cpuprofile=cpu.out
//go tool pprof cpu.out

// (pprof) top
// flat  flat%   sum%        cum   cum%
//     51.29s 99.94% 99.94%     51.29s 99.94%  github.com/dazulu4/testing.Fibonacci
//          0     0% 99.94%     51.29s 99.94%  github.com/dazulu4/testing.TestFib
//          0     0% 99.94%     51.29s 99.94%  testing.tRunner

// (pprof) list Fibonacci
// Total: 51.32s
// ROUTINE ======================== github.com/dazulu4/testing.Fibonacci in /Users/dzuluaga/Proyectos/golang-in-practice/platzi/03_intermediate/17_testing/main.go
//     51.29s        91s (flat, cum) 177.32% of Total
//     28.12s     28.12s     14:func Fibonacci(n int) int {
//      400ms      400ms     15:   if n <= 1 {
//     11.19s     11.19s     16:           return n
//          .          .     17:   }
//     11.58s     51.29s     18:   return Fibonacci(n-1) + Fibonacci(n-2)
//          .          .     19:}

//web
//Reporte HTML

//pdf
//Reporte archivo PDF
