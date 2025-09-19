package main 

func main(){
	IsPalindrome(121)
}

func IsPalindrome(x int) bool {
    strx:= Itoa(x)
    var newx string
    if x < 0{
        newx = "-"
        for i:=0 ; i < len(strx) ; i++{
            newx += string(strx[i])
        }
    }else{

   for i:= len(strx)-1 ; i >=0 ; i--{
	newx += string(strx[i])
   }
    }
  
    if newx == strx{
        return true
    }else{
        return false
    }

    
}
func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
	}
	digits := []byte{}
	for n > 0 {
		d := n % 10
		digits = append(digits, byte('0'+d))
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
    return string(digits)
}