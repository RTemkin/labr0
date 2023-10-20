package main

import (
	"fmt"
	"time"
)

func ab(a int, c chan []int) {
	nums := make([]int, a)
	for i := 0; i < len(nums); i++ {
		nums[0] = 1
		nums[i] = i + 1
	}
	c <- nums
}

func dubl(nums []int, c1 chan []int) {
	nums1 := make([]int, len(nums))
	for v := range nums {
		nums1[v] = nums[v] * 2
	}
	c1 <- nums1
}

func main() {
	start := time.Now()
	a := 1000000
	c := make(chan []int)
	c1 := make(chan []int)
	go ab(a, c)
	nums := <-c

	n := 0
	i := 250000
	m := []int{}
	for i <= len(nums) {
		go dubl(nums[n:i], c1)
		n = i
		i = i + 250000
		x := <-c1
		m = append(m, x...)

	}

	t := time.Since(start)

	fmt.Println(m)
	fmt.Println(t)
}
