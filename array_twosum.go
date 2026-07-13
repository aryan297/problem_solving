package main

import "fmt"

func twoSum(arr []int , target int) []int {

	hashMap := make(map[int]int)

	for i:=0;i<len(arr);i++{

		comp := target -arr[i]

		if j , ok := hashMap[comp];ok{
			return []int{i,j}
		}

		hashMap[arr[i]]=i


	}


	return []int{}


}

func main(){

	nums := []int{2, 7, 11, 15}

	target := 9

	fmt.Println(twoSum(nums,target))
}