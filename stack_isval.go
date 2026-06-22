package main

import "fmt" 

func isValid(str string) bool{

	maps := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	stack :=[]byte{}

	for _, r := range str{
		 fmt.Println(r,"ddd")
		  ch:=byte(r)

		  fmt.Println(ch,"sss")

		if ch=='(' || ch=='{' || ch== '[' {

			stack=append(stack,ch)

		} else{

			if len(stack)==0 || stack[len(stack)-1]!= maps[ch]{
				return false
			}

			stack=stack[:len(stack)-1]
		}
	}

	return len(stack)==0


}


func main(){

	str :="{}[]"

	k := isValid(str)

	fmt.Println(k )
}