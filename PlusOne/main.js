 function plusOne (digits) {
    for(let i = digits.length-1 ; i >=0 ; i--){
        if (digits[i]!== 9){
            digits[i] = digits[i]+1
            break
        }else if(digits[i] === 9 && i !==0){
            digits[i] = 0
        }else if(digits[i] === 9 && i === 0){
            digits[i]=0            
            return [1,...digits]
        }
    }
    return digits
};
console.log(plusOne([9,9,9]));
