package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}

func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Reversing the number
	original := x
	reversedNum := 0
	for x != 0 {
		digit := x % 10
		reversedNum = reversedNum*10 + digit
		x = x / 10
	}

	// Check if the original number is equal to the reversed number
	return original == reversedNum
}

// Algorithm
// Reverse the number without using a string
// 1. Check if the number is negative
// 2. Initialize a variable to store the original number
// 3. Initialize a variable to store the reversed number
// 4. Loop through the number until it becomes 0
// 5. Get the last digit of the number
// 6. Add the digit to the reversed number
