package main

import (
	"log"
)

func main() {
	number := 12221

	log.Println("Palindrome Number", number)

	log.Println(isPalindrome(number))
}

func isPalindrome(x int) bool {
	reverseNumber := reverseNumber(x)

	if x == reverseNumber {
		return true
	} else {
		return false
	}
}

func reverseNumber(num int) int {
	reverse := 0

	for num > 0 {
		remainder := num % 10
		reverse = (reverse * 10) + remainder
		num /= 10
	}

	return reverse
}
