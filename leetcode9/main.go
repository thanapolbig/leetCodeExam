package main

import "log"

func main() {
	result := isPalindrome(121)
	log.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := 0
	value := x
	for value != 0 {
		rev = rev*10 + value%10
		value /= 10
	}
	return rev == x
}
