package main


import "fmt"

func reverseArray(s string){
	runes:=[]rune(s)

 left:=0;
 right:=len(runes)-1


 for left < right{
 	runes[left],runes[right]=runes[right],runes[left]
 	left++
 	right--


 }

 fmt.Println(string(runes))

}


func main(){

	str:="ram"
	reverseArray(str)
}