package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	fmt.Print(sum)
}
