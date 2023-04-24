package main

import "fmt"

func main() {
	n := 1000000 // number of iterations
	sum := 0.0
	sign := 1.0
	denom := 1.0
	for i := 0; i < n; i++ {
		sum += sign / denom
		sign *= -1
		denom += 2
	}
	pi := 4 * sum
	fmt.Println("π is approximately:", pi)
}
