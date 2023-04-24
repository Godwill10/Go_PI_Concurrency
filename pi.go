package main

import (
    "fmt"
    "math/rand"
)

func main() {
    n := 1000000 // number of iterations
    count := 0   // count of points inside the circle
    for i := 0; i < n; i++ {
        x := rand.Float64()
        y := rand.Float64()
        if x*x+y*y < 1 {
            count++
        }
    }
    pi := 4.0 * float64(count) / float64(n)
    fmt.Println("π is approximately:", pi)
}
