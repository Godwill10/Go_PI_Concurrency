**Concurrency in Go**

**Name: Godwill Afolabi**

**Purpose**

This is a simple Go program that estimates the value of pi using the stochastic method. The program defines an array of iteration values, and then spawns a goroutine for each value in the array. Each goroutine performs the stochastic method with the given number of iterations to estimate pi, calculates the difference between the estimated pi and the actual pi, and reports the results.


**Organization**

The program is structured as follows:

    1. Define an array of iteration values.

    2. Spawn a goroutine for each iteration value in the array.

    3. Perform the stochastic method with the given number of iterations to estimate pi.

    4. Calculate the difference between the estimated pi and the actual pi.

    5. Report the results.

    The program uses the sync.WaitGroup type to synchronize the goroutines, and the defer statement to decrement the wait group counter when a goroutine completes its execution. The rand.Float64() function is used to generate random numbers for the stochastic method.


**How to run/use**

To install and run the program, follow these steps:

1. Install Go on your computer. You can download it from the official website: https://golang.org/doc/install

2. Clone this repository or download the source code to your computer.

3. Open a terminal or command prompt and navigate to the directory containing the source code.

4. Run the following command to compile and run the program:

    go run pi.go

5. The program will run and output an estimated value of pi.

