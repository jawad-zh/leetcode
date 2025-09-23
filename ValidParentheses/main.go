package main

import "fmt"

func main() {
	Result := isValid("({})")
	fmt.Println(Result)
}

func isValid(s string) bool {
	var RoundBrackets int
	var SquareBrackets int
	var CurlyBrackets int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			RoundBrackets++
			}else if s[i] == ')' {
				RoundBrackets--
				if (CurlyBrackets != 0 || SquareBrackets != 0) && RoundBrackets!=0 {
				return false
			}
			} else if s[i] == '{' {
				CurlyBrackets++
			} else if s[i] == '}' {
				CurlyBrackets--
					if (RoundBrackets != 0 || SquareBrackets != 0) && CurlyBrackets !=0 {
					return false
				}
			} else if s[i] == '[' {
				
				SquareBrackets++
			} else if s[i] == ']' {
				SquareBrackets--
				if (RoundBrackets != 0 || CurlyBrackets != 0) && SquareBrackets!=0 {
					return false
				}
			}
		}
		if RoundBrackets == 0 && SquareBrackets == 0 && CurlyBrackets == 0 {
			return true
		}
		return false
	}
