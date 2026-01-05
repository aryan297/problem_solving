package main

import "fmt"


func twoSum(arr []int , target int) []int{ 

	hashMap:=make(map[int]int)

	for i:=0;i<len(arr);i++{

		comp:=target-arr[i]

		if idx,found:=hashMap[comp]; found{
			return []int{arr[i],arr[idx]}
		}
         hashMap[arr[i]]=i

	}
  return []int{}
}


func main(){
	//nums = [2, 7, 11, 15], target = 9

	arr:=[]int{2, 7, 11, 15}
	target:=9

	fmt.Println(twoSum(arr,target))
}