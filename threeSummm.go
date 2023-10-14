package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if int(math.Abs(float64(sum-target))) < int(math.Abs(float64(res-target))) {
				res = sum

			} else if sum > target {
				k--
			} else if sum < target {
				j++
			}

		}
	}
	return res
}

func main() {
	nums := []int{-1, 2, -5, 3, -2, 1, 4, 8}
	target := 16

	fmt.Println(threeSumClosest(nums, target))

}
