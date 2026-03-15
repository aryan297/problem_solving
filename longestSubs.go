package main

import "fmt"

func subStr( s string){

	hashMap := map[byte]bool{}

	left:=0

	minLen:=0
	for right:=0;right<len(s);right++{



		for hashMap[s[right]]{

			delete(hashMap,s[left])
			left++
		}

		hashMap[s[right]]=true

		if right-left+1 > minLen{
			minLen=right-left+1
		}
	}

fmt.Println(minLen)

}

func main(){
	s := "abcabcbb"

	subStr(s)
}