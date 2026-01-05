package main


import "fmt"


func unique(arr []int){

	hashMap:=make(map[int]int)

	for i:=0;i<len(arr);i++{
		hashMap[arr[i]]++
	}


   for k :=range hashMap{

   	if hashMap[k]==1{
   		fmt.Println(k)
   	}
   }





}



func main(){

	arr :=[]int{1,1,2,3,4,3,4}

	unique(arr)
}

