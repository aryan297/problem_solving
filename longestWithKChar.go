package  main

import "fmt"

func longestChar(s string , k int){

	left:=0
	hashMap:=make(map[byte]int)

	maxLen:=0

	for right:=0;right<len(s);right++{

		hashMap[s[right]]++


		for len(hashMap)>k{
			hashMap[s[left]]--

			if hashMap[s[left]]==0{

				delete(hashMap,s[left])
			}

			left++

		}

		if right-left+1>maxLen{
			maxLen=right-left+1
		}
	}

	fmt.Println(maxLen)



}






func main(){
	s :="eceba"
	k :=2

	longestChar(s,k)
}