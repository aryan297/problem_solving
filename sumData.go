package main

import "fmt"


func sums(arr []int , k int){

	sumMap :=make(map[int]int)

	sum , maxLen:=0,0

	for i,val := range arr{

		sum +=val;

		if sum==k{
			maxLen=i+1
		}

		if idx, found := sumMap[sum-k]; found{

			if i-idx>maxLen{
				maxLen=i-idx
			}
		}

		// store first occurrence of sum
		if _, exists := sumMap[sum]; !exists {
			sumMap[sum] = i
		}

	
	}


 fmt.Println(maxLen)

}


func main(){
	arr :=[]int{2,3,5,1,9}
	sums(arr,10)
}