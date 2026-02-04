package main

import "fmt"

func subArray(arr []int, k int)  int{

	hashMap:=make(map[int]int)

	count:=0

	prefixSum :=0
	hashMap[0]=1

	for i:=0;i<len(arr);i++{
      prefixSum+=arr[i]


		if idx, exist := hashMap[prefixSum-k];exist{

			count +=idx
		}
       
      hashMap[prefixSum]++
	}

	return count



}


func main(){
	arr :=[]int{1,2,3} // {1,2} ,{3} output 2
	value:=subArray(arr,3)
	fmt.Println(value)
}