package main

import "fmt"

func main() {
	var Names []string

	Names = append(Names, "jaw")
	Names = append(Names, "jawa")
	Names = append(Names, "jawd")
	longestCommonPrefix(Names)
}

func longestCommonPrefix(strs []string) string {
	var Result string
	var Same bool
	var Rune string

	if len(strs) > 1 {
		smallLenght := strs[0]
		for _, str := range strs {
			if len(str) < len(smallLenght) {
				smallLenght = str
			}
		}

		j := 0
		for j < len(smallLenght) {
			Same = true
			for i := 1; i < len(strs); i++ {
				if strs[i][j] != strs[0][j] {
					Same = false
					break
				}
			}

			if Same {
				Rune = string(strs[0][j])
				Result += Rune
				j++
			} else {
				break
			}
		}
	} else if len(strs) == 1 {
		return strs[0]
	}

	fmt.Println(Result)
	return Result
}
