package main

import "fmt"

func findOdd(s string){

   max:=0

for i:=len(s)-1 ;i>=0;i--{
	digit:=s[i]-'0'
	if digit%2==1{
		res:=s[i];
        if int(res)>max{
        	max=int(res)
        }

		fmt.Println(string(max))

	    
	}
}
}




func main(){

s:="5467"

findOdd(s)


}