package main

import "fmt"

func main() {
	fmt.Println(addBinary("10101","110"))	
}

func addBinary(a string, b string) string {
	counter := 0;
	carry := 0;
 
	var strBuilder string
	for counter < len(a) && counter < len(b){

		x := a[len(a)-1- counter];
		y := b[len(b)-1- counter];
            if carry == 1 {
                if x=='1' && y=='1' {
        		strBuilder += string("1");
                }else if x=='0' && y=='0' {
					strBuilder += string("1");
                    carry = 0;
                }else{
					strBuilder += string("0");
                }       
            }else{
                if x=='1' && y=='1' {
					strBuilder += string("0");
                    carry=1;
                }else if x=='0' && y=='0' {
					strBuilder += string("0");
                }else{
					strBuilder += string("1");
                }
            }   
            counter++;
	}

	for counter < len(a){
		x := a[len(a) - 1 - counter];
		if carry==1 {
		  if x=='0' {
			strBuilder += string("1");
			  carry = 0;
			}else{
        		strBuilder += string("0");
				carry=1;
			}    
		}else{
			strBuilder += string(x);		}
		counter++;
	}
	for counter < len(b){
		y := b[len(b) - 1 - counter];
		if carry==1 {
			if y=='0' {
        		strBuilder += string("1");
				carry = 0;
			}else{
        		strBuilder += string("0");
				carry=1;
			}
		}else{
			strBuilder += string(y);		}
		counter++;
	}
	
	if carry==1 {
		strBuilder += string("1");
	}
	return reverse(strBuilder);
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func min(a,b int) int{
	if a > b{
		return a
	}
	return b
}
