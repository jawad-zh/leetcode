package main

import "fmt"

func main() {
	longestPalindrome("rotavtolevel")
}

func longestPalindrome(s string) string {
	sRune := []rune(s)
	var comparS []rune
	var word string
	var words []string
	for i := len(sRune) - 1; i >= 0; i-- {
		comparS = append(comparS, sRune[i])
	}
	k := 0
	for k < len(comparS) {
		for i := 0; i < len(sRune) ; i++ {
			fmt.Println("k",k)
			fmt.Println("i",i)
			if comparS[k] == sRune[i] {
				fmt.Println(string(comparS[k]))
				word += string(comparS[k])
				k++
			} else {
				if word != "" {
					words = append(words, word)
					word = ""
					break
				}
			}
		}
		if word !=""{
			words = append(words, word)
			word = ""
		} 
	}
	if word != ""{
		words = append(words, word)
	}
	

	return ""
}
