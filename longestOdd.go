package main


import "fmt"

func subString(s string) string{

	for i:=len(s)-1; i>0;i--{

		digit:= s[i]-'0'

		if digit%2==1{

			return s[:i+1]

		}
	}

	return ""

}



func main(){
	s:="532"

	fmt.Println(subString(s))
}