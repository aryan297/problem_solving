package main

import "fmt"

func longestSubString(s string){
	left:=0
	hashMap:=make(map[byte]bool)
    maxLen:=0

    for right:=0;right<len(s);right++{

    	for hashMap[s[right]]{
    		delete(hashMap,s[left])
    		left++

    	}

    	hashMap[s[right]]=true

    	if right-left+1>maxLen{
    		maxLen=right-left+1
    	}

    }

    fmt.Println(maxLen)


}


func main(){
	s:="abcabcbb"
	longestSubString(s)
}