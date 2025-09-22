package main

import "fmt"

func main() {
	// RomanToInt("III")
	RomanToInt("MCMXCVI")
}

func RomanToInt(s string) int {
	var Result int
	var slice []string
	var sliceint []int
	key:= map[string]int{
		"I":1,
		"V":5,
		"X":10,	
		"L":50,	
		"C":100,	
		"D":500,	
		"M":1000,	
		"two":2,
		"three":3,
	}
	for _,r:= range s{
		slice = append(slice, string(r))
	}
	for i:=0 ; i < len(slice) ; i++{
		if slice[i]== "I" && i!=0{
			if slice[i-1] == "I"{
				slice [i] = ""
				slice[i-1]= "two"	
				i--
			}else if slice[i-1]== "two"{
				slice[i]= ""
				slice[i-1]="three"
				i--
			}
			slice = clean(slice)
		}
		
	}
	for i:=0 ; i < len(slice) ; i++{
		if value,ok:=key[slice[i]] ; ok{
			sliceint = append(sliceint, value)
		}
	}
	for j:=0 ; j < len(sliceint)  ; j++{
		if len(sliceint)>1 {
		if j!=0 && sliceint[j] > sliceint[j-1]{
			sliceint[j]= sliceint[j] - sliceint[j-1]
			sliceint[j-1]=0
		}
		
		}
	}
	for i:=0 ; i < len(sliceint) ; i++{
		Result += sliceint[i]
	}
	fmt.Println(Result)
	return Result
	
}

func clean(slice []string)([]string){
	var Result []string 
	for i:=0 ; i < len(slice); i++{
		if slice[i]!= ""{
			Result = append(Result,slice[i])
		}
	}
	
	return Result
}