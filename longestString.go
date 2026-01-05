package main

import ("fmt"
	    "sort"
	    "strings")


func longestSub(str string){

	seen:=make(map[rune]int)
	i :=0
	maxLen:=0

	for right,ch := range str{

		if idx, ok :=seen[ch];ok && idx >=i{
			i=i+1
		}

		seen[ch]=right

		if right-i+1>maxLen{
			maxLen=right-i+1

		}
	}
   
   fmt.Println(maxLen)


}

func groupOfAnagram(str []string) {
     anagramMap:=make(map[string][]string)
     res :=[][]string{}


     for _ , word :=range str {
     	sortedString :=sortData(word)
     
        anagramMap[sortedString]=append(anagramMap[sortedString],word)

     }

     for _ ,group := range anagramMap{
     	res=append(res,group)
     }
	fmt.Println(res)
}

func sortData(str string) string{
    chars:=strings.Split(str,"")
    sort.Strings(chars)
   // fmt.Println(chars)
    return strings.Join(chars,"")
}






func main(){
	str:="abcabcbb"
	  input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	longestSub(str)
	groupOfAnagram(input)

}