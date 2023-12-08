package main

// func mergeSlise(nums1, nums2 []int) []int {
// 	if nums1 == nil {
// 		return nums2
// 	}
// 	if nums2 == nil {
// 		return nums1
// 	}
// 	size, i, j := len(nums1)+len(nums2), 0, 0
// 	nums := make([]int, size)

// 	for k := 0; k < size; k++ {
// 		if i > len(nums1)-1 && j <= len(nums2)-1 {
// 			nums[k] = nums2[j]
// 			j++
// 		} else if i <= len(nums1)-1 && j > len(nums2)-1 {
// 			nums[k] = nums1[i]
// 			i++
// 		} else if nums1[i] < nums2[j] {
// 			nums[k] = nums1[i]
// 			i++
// 		} else {
// 			nums[k] = nums2[j]
// 			j++
// 		}
// 	}
// 	return nums
// }

// func sort1234(nums []int) []int {
// 	if len(nums) < 2 {
// 		return nums
// 	}
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] > nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 		}
// 	}
// 	return nums
// }

// func main() {
// 	nums1 := []int{1, 18, 4, 2, 4, 5, 25, 19, 8, 4, 5, 17}
// 	nums2 := []int{0, 6, 2, 3, 1, 6, 7, 10, 14, 15, 16, 4, 11}
// 	a := mergeSlise(nums1, nums2)

// 	fmt.Println(sort1234(a))
// }

type ListNode struct {
	value int
	Next  *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.value < list2.value {
			current.Next = list1
			list1 = list1.Next
		} else if list1.value >= list2.value {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}
	return result.Next
}

func main() {

}
