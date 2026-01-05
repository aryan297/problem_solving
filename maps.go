package main

import "fmt"


func interSection(arr ,arr1 []int){
   hashMap:=make(map[int]int)
   result:=[]int{}

   for i:=0;i<len(arr);i++{
   	  hashMap[arr[i]]++
   }

   for i:=0;i<len(arr1);i++{
   	  if hashMap[arr1[i]]>0{
   	  	result=append(result,arr1[i])
   	  hashMap[arr1[i]]--

   	  }

   }


   fmt.Println(result)
}


func main(){

	arr:=[]int{1,3,3,1}
	arr1:=[]int{1,2,3,3}

interSection(arr,arr1)

}