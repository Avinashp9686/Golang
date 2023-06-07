package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Enter the value of n")
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
}
