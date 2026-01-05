package main

import "fmt"

func longest(s string){
	hashMap := make(map[rune]int)
        start, maxLen := 0, 0
	for i,ch:=range s{

		if idx,found:=hashMap[ch];found && idx >=start{
			start=idx+1
		}

		hashMap[ch]=i

		if i-start+1 >maxLen{
			maxLen=i-start+1
		}


	}

	fmt.Println(maxLen)
}





func main(){

	s:="abcb"

	longest(s)

}