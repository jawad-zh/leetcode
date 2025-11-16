package main

import "fmt"

func main() {
	 test := []int{9,9}
 plusOne(test)
}

func plusOne(digits []int) []int {
	var stop = 1
	var tracker = len(digits) - 1
	var temp []int
	for stop > 0 {
		if len(digits) > 1 {
			if digits[tracker] != 9 {
				digits[tracker]++
				stop--
			} else {
			if tracker == 0{
			digits[0] = 0
			temp= append(temp,digits...)
			digits =nil 
			digits = append(digits, 1)
			digits = append(digits, temp...)
			break
			
		}
				digits[tracker] = 0
				tracker--
			}
		} else {
			if digits[tracker] != 9 {
				digits[tracker]++
				stop--
			} else {
				digits = nil
				digits = append(digits, 1)
				digits = append(digits, 0)
				break
			}
		}
	}
	return  digits
}
