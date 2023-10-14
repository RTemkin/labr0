package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			start, end := mid, mid
			for start > 0 && nums[start-1] == target {
				start--
			}
			for end < len(nums)-1 && nums[end+1] == target {
				end++
			}
			return []int{start, end}
		}
	}
	return []int{-1, -1}
}

func main() {

	nums := []int{9, 9, 7, 7, 7, 5, 5, 3, 2, 1, 0}
	target := 7
	sort.Ints(nums)
	result := searchRange(nums, target)

	fmt.Println(result)
	fmt.Println(nums)
	fmt.Println(len(nums))

}
