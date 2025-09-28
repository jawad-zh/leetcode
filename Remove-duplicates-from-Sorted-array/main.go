package main

import "fmt"

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	RemoveDuplicates(test)
}

func RemoveDuplicates(nums []int) int {
	i := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)
	return i + 1
}
