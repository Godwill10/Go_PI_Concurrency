**Concurrency in Go**

**Name: Godwill Afolabi**

**Purpose**

This is a simple Go program that estimates the value of pi using the stochastic method method. The program defines an array of iteration values, and then spawns a goroutine for each value in the array. Each goroutine performs the stochastic method with the given number of iterations to estimate pi, calculates the difference between the estimated pi and the actual pi, and reports the results.


**Organization**




**How to run/use**

To install and run the program, follow these steps:

1. Install Go on your computer. You can download it from the official website: https://golang.org/doc/install

2. Clone this repository or download the source code to your computer.

3. Open a terminal or command prompt and navigate to the directory containing the source code.

4. Run the following command to compile and run the program:

    go run pi.go

5. The program will run and output an estimated value of pi.

