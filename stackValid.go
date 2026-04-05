package main

import "fmt"


func checkValid(s string) bool{
	stack :=[]rune{}

	maps:= map[rune]rune{
		'}':'{',
		')':'(',
		']': '[',
	}


	for _, ch :=range s {

		if open,exist:=maps[ch];exist{

			if len(stack)==0 || stack[len(stack)-1]!=open{
				return false
			}

			stack=stack[:len(stack)-1]
		} else{

			stack=append(stack,ch)
		}
	}

	return true

}


func main(){
	s := "{()}"
	k:=checkValid(s)
	fmt.Println(k)
}