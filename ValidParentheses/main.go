package main

import "fmt"

func main() {
	Result := isValid("({)}")
	fmt.Println(Result)
}

func isValid(s string) bool {
	var Result bool
	var RoundOver bool 
	var CurlyOver bool 
	var SquarOver bool 
	var RoundIndex int 
	var CurlyIndex int 
	var SquarIndex int 
	for i:=0 ; i < len(s) ; i++{
		for j:=i+1 ; j < len(s) ; j++{
		if s[i]== '(' {
			RoundOver = true 
			RoundIndex = i
			 if s[j]== ')'{
			if (CurlyOver == true || SquarOver == true ) && (RoundIndex < CurlyIndex || RoundIndex < SquarIndex)  {
				fmt.Println(Result)
				return false
			}
		}
		}else if s[i]== '{' {
			CurlyOver = true 
			CurlyIndex = i
			 if s[j]== '}'{
			if (RoundOver == true || SquarOver == true ) && (CurlyIndex < RoundIndex || CurlyIndex < SquarIndex)  {
				fmt.Println(Result)
				return false
			}
		}
		}else if s[i]== '[' {
			SquarOver = true 
			SquarIndex = i
			 if s[j]== ']'{
			if (CurlyOver == true || RoundOver == true ) && (SquarIndex < CurlyIndex || SquarIndex < RoundIndex)  {
				fmt.Println(Result)
				return false
			}
		}
		}
	}
	}
	// fmt.Println(true)
	return true
}