package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	nums1 := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {

				nums1 = append(nums1, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j-1] == nums[j] {

					j++
				}
			}
		}

	}

	return nums1
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}
