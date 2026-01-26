package main

import "fmt"


func subStr(s string){
	charMap := make(map[byte]int)
	left:=0
	maxLen:=0

	for right:=0;right<len(s);right++{

		if idx, exist :=charMap[s[right]];exist && idx >=left{
			left=idx+1
		}

	    charMap[s[right]]=right
	    maxLen = max(maxLen, right - left + 1)

	}

	fmt.Println(maxLen)
}



func max(a , b int)int{
	if a>b{
		return a
	}

	return b
}

func main(){
	str :="abcabcbb"
	subStr(str)
}