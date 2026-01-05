package main

import "fmt"

func minimumLength(s string , t string){

	left:=0
	found:=0
	 need :=make(map[byte]int)
	 window:=make(map[byte]int)
	 required:=len(t)
	 minLen:=len(s)+1
	 start:=0

	 for i:=0;i<len(t);i++{
	 	need[t[i]]++
	 }

	 for right:=0;right<len(s);right++{
	 	window[s[right]]++
	 

	 if count,ok:= need[s[right]];ok && window[s[right]]==count{
	 	found++
	 }

	 for found==required{
       
        if right-left+1 < minLen{
        	minLen=right-left+1
        	start=left
        }

        window[s[left]]--

        if count,ok:= need[s[left]];ok && window[s[left]]<count{
        	found--
        }

        left++

	 }


}
	 fmt.Println(s[start:start+minLen])

}

func main(){
s := "ADOBECODEBANC"
t := "ABC"

minimumLength(s,t)
}