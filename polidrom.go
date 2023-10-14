package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {

	x := 100
	result := isPalindrome(x)

	fmt.Println(result)

}
