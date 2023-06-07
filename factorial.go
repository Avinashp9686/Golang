package main

import "fmt"

func main() {
	var n int
	factorial := 1
	fmt.Println("Enter the numbers")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}
	fmt.Print(factorial)
}
