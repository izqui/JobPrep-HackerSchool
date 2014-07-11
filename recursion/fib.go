package main

import "fmt"

func main() {

	go fmt.Println("mem", mem_fib(100000))
	go fmt.Println("naive", naive_fib(100))

}

func naive_fib(n int) int {

	if n > 1 {

		return naive_fib(n-1) + naive_fib(n-2)

	} else {

		return n
	}
}

func mem_fib(n int) int {

	var fib func(a, b, n int) int
	fib = func(a, b, n int) int {

		if n == 0 {

			return a
		} else {

			return fib(b, a+b, n-1)
		}
	}

	return fib(0, 1, n)
}
