package main

import "fmt"


func isAnagram(str , str1 string) bool{

   hashMap:=make(map[rune]int)

   if len(str)!= len(str1){
   	return false 
   }

   for i:=0;i<len(str);i++{
   	hashMap[ rune(str[i])]++
   }


   for i:=0;i<len(str1);i++{
   	   hashMap[rune(str1[i])]--;
   	   if hashMap[rune(str1[i])] < 0 {
   	   	return false

   	   }
   }

 return true


}



func main(){
	str:="listen"
	str1:="silenj"

	fmt.Println(isAnagram(str,str1))
}