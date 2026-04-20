package main

import "fmt"

func validAna(s string , t string) bool{

	if len(s) != len(t){
		return false
	}

	var count [26]int


	for i:=0;i<len(s);i++{
		count[s[i]-'a']++
		count[t[i]-'a']--

	}


	for _ , ch := range count{

		if ch!=0{
			return false
		}
	}

	return true;


}


func main(){
	s := "anagram"
	t := "gramana"

k :=validAna(s,t)

fmt.Println(k)
}