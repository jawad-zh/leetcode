package main

import "fmt"

func main() {
	RomanToInt("LVIII")
}

func RomanToInt(s string) int {
	var slqueue []int
	var queue int
	var one int
	first := true
	var Result int
	var other bool
	Map := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i := 0; i < len(s); i++ {
		value,ok:= Map[string(s[i])]
		if ok {
			slqueue = append(slqueue, value)
		}
	}
	for i:=0 ; i < len(slqueue) ; i++{
		fmt.Println(slqueue[i])
		if int(slqueue[i]) == 1{
			one+= 1
			fmt.Println("one:",one)
		}else{
			other = true
			if one >= 2 && one <=3{
				Result += one
				fmt.Println("oneenterd")
			
			}
			one=0

				if first{
					Result+=slqueue[i] 
					first = false
				}
				if i!= 0 && i%2==0{
					queue = slqueue[i]-slqueue[i-1]
					Result += queue
					queue = 0
				}
			
		}
		}
		if !other{
			Result = one
		}
		fmt.Println(Result)
return Result
}