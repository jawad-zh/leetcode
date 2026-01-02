package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 1,3,1,4,2,3,4,1,3,5,1,4,4)
	fmt.Println(repeatedNTimes(arr))
}
func repeatedNTimes(nums []int) int {
	NumbersTimes := make(map[int]int)
    for _,num := range nums{
		NumbersTimes[num]++
		if NumbersTimes[num]>1{
			return num
		}
	}
	return -1
}