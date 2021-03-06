package main

import "fmt"

func fibo(first, second int64) int64 {
	return first + second
}

func main() {
	x, y := int64(1), int64(1)
	var dummy, sum int64
	for y <= 4000000 {
		dummy = fibo(int64(x), int64(y))
		if dummy%2 == 0 {
			sum += dummy
		}
		x = y
		y = dummy
	}
	fmt.Println(sum)
}

// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
