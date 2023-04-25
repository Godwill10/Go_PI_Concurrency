package main

import (
    "fmt"
    "math/rand"
	"time"
)

func main() {
	// Define the number of iterations for each experiment
	iterations := []int{10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000}

	// Loop through each iteration value and perform the experiment
    for i := 0; i < len(iterations); i++ {
        start := time.Now()

        // Perform the Monte Carlo method to estimate pi
        count := 0
        for i := 0; i < n; i++ {
            x := rand.Float64()*2 - 1
            y := rand.Float64()*2 - 1
            if x*x+y*y < 1 {
                count++
            }
        }
        pi := 4.0 * float64(count) / float64(n)

		// Calculate the delta between the estimated pi and the actual pi
		delta := pi - 3.14159265358979323846

		// Report the results
		fmt.Printf("Number of iterations: ", n)
		fmt.Printf("Estimated pi: ", pi)
		fmt.Printf("Delta from true pi: ", delta)
		fmt.Printf("Time taken: ", time.Since(start))
    }
}
