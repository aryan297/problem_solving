package main

import "fmt"



func countLeast(arr []int){

	hashMap:=make(map[int]int)
	result:=[]int{}

	for i:=0;i<len(arr);i++{
		hashMap[arr[i]]++
	}


  for i:=0;i<len(arr);i++{
  	if hashMap[arr[i]]==1{
  		result=append(result,arr[i])
  	}
  }
   fmt.Println(result)


}

func main(){

arr := []int{7, 3, 5, 3, 7, 9, 5}
countLeast(arr)

}