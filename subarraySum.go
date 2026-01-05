package main

import "fmt"

func subarraySum(arr []int , k int) int{
	count:=0
	prefixSum:=0
	subArr:=map[int]int{0:1}

	for _ , num:= range arr{
		prefixSum += num

		if subArr[prefixSum-k]>0{
			fmt.Println(prefixSum-k)
			count+=subArr[prefixSum-k]
		}
		subArr[prefixSum]+=1
	}

	return count;

}



func main(){
	arr :=[]int{1,2,3}
	k:=3
	fmt.Println(subarraySum(arr,k))
}