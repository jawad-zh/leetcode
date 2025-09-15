package main

import "fmt"

func main() {
	test := "hello world khoro"

	brk:= "adm"
	CanBeTypedWords(test,brk)

}


func canBeTypedWords(text string, brokenLetters string) int {
var word string 
var words []string 
var count int
 FinishWord:= false
for i:=0 ; i < len(text) ; i++{
    if string(text[i]) != " "{
        word += string(text[i])
    }else{
        if word != ""{
            words = append(words,word)
            word = ""
        }
    }
}
if word != ""{
    words = append(words,word)
    word = ""
}
if brokenLetters == ""{
 return len(words)   
}
for j:= 0 ; j < len(words);j++{
    for k:=0 ; k < len(brokenLetters) ; k++{
      if  check(words[j],rune(brokenLetters[k])) && k == len(brokenLetters)-1 {
        FinishWord = true
        
      }else if ! check(words[j],rune(brokenLetters[k])){
        if len(words)>1 && j!= len(words)-1{
        break
        }
      }
    }
    if FinishWord{
        count++
        FinishWord  = false
    }
}
return count
}
func check(s string , r rune)bool{
for _,ru:= range s {
    if ru == r{
        return false
    }
}
return true
}