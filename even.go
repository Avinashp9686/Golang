package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("enter the value of n")
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("sum:", sum)

}
