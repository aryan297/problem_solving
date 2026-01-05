
package main

import "fmt "


func anagram(s1,s2 string){
	char:=make(map[rune]int)

	for _,str :=range s1{
		char[str]++
	}


	for _, str:=range s2{
		char[str]--
		if char[str]<0{
			fmt.Println(false)
			return
		}
	}

	fmt.Println(true)
	return
}


func main(){

	s1:="listen"
	s2:="silent"

	anagram(s1,s2)
}