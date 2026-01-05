package main

import "fmt"


func twoSum(arr []int, target int){

	hashMap := make(map[int]int)

	for i:=0;i<len(arr);i++{

		if idx, ok := hashMap[target-arr[i]];ok{

			 fmt.Println(idx,i)
		}

		hashMap[arr[i]]=i
	}
}

func main(){
	arr :=[]int{2, 11, 15,7}

	target:=9

	twoSum(arr,target)
}