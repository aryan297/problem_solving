package main

import "fmt"


func lengthSub(s string){

	left :=0
	best :=0

	lastSeen := make(map[byte]int)

	for right:=0; right < len(s) ;right++{

		if idx, ok := lastSeen[s[right]]; ok  && idx >=left{
			left=idx+1
		}

		lastSeen[s[right]]=right

		if right-left+1 >best{
			best=right-left+1
		}



	}

	fmt.Println(best)


}


func main(){
	s := "abcabcbb"
	lengthSub(s)


}