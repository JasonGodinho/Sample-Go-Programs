package main

import "fmt"

func RecursuveFibonacci(fibonacci_num int) int {
	if fibonacci_num == 1 {
		return 1
	} else if fibonacci_num == 0 {
		return 0
	} else {
		return RecursuveFibonacci(fibonacci_num-1) + RecursuveFibonacci(fibonacci_num-2)
	}
}

func main() {
	var fibonacci_num int
	fmt.Println("Please enter the fibonacci number")
	fmt.Scanln(&fibonacci_num)

	if !(fibonacci_num < 0) {
		fibonacciSum := RecursuveFibonacci(fibonacci_num)
		fmt.Println(fibonacciSum)
	} else {
		fmt.Println("Please enter a valid number from 0 onwards")
	}
}
