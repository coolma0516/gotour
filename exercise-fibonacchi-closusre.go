package main

import "fmt"

// fibonacchi is a function that returns
// a function that returns an int

func fibonacchi() func() int {
	first := 0
	second := 1
	return func() int {
		res := first
		first, second = second, first+second
		return res
	}
}

func main() {
	f := fibonacchi()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
