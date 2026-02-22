package main

import "fmt"

func longest(s string){
	hashmap :=make(map[byte]int)

	maxLen :=0
	left:=0

	for right :=0;right<len(s);right++{

		hashmap[s[right]]++

		for hashmap[s[right]] > 1{
			hashmap[s[left]]--
			left++
		}

		if right-left+1>maxLen{
          maxLen=right-left+1
		}
	}


fmt.Println(maxLen)

}


func main(){

   s := "abcabcbb"
   longest(s)


}