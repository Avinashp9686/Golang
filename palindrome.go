package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
func largestPalindromeProduct() (int, int, int) {
	largestPalindrome := 0
	var multiplicand1, multiplicand2 int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if product < largestPalindrome {
				break
			}
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
				multiplicand1 = i
				multiplicand2 = j
			}
		}
	}
	return largestPalindrome, multiplicand1, multiplicand2
}
func main() {
	result, multiplicand1, multiplicand2 := largestPalindromeProduct()
	fmt.Printf("The largest palindrome product is:%d\n", result)
	fmt.Printf("The multiplicand are:%d and %d\n", multiplicand1, multiplicand2)
}
