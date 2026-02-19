package main

import "fmt"

func twoSum(arr []int , target int){

	hashMap:=make(map[int]int)

	for i:=0;i<len(arr);i++{

		diff := target-arr[i]

		if idx,ok := hashMap[diff];ok{
			fmt.Println(idx,i)
		}
		hashMap[arr[i]]=i

	}

}


func main(){
	arr :=[]int{2,3,4,5,4}

	twoSum(arr,9)
}