package main

import (
	"fmt"
)

func reverse(x int) int {
	result := 0

	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}

	if int(int32(result)) != result {
		return 0
	}
	return result
}

func main() {

	x := -32854
	fmt.Println(reverse(x))
}
