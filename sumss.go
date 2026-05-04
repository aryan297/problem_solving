package main

import "fmt"

func twoSum(arr []int , k int) []int{
	hashMap:=make(map[int]int)

	for i,num := range arr{

		complement := k-arr[i]

		if val , ok:=hashMap[complement];ok{

			return []int{val,i}
		}

		hashMap[num]=i

	}

	return []int{}
}



func main() {

	arr :=[]int{2, 7, 11, 15}
   
   n:=twoSum(arr,9)

   fmt.Println(n)


}