package main

import "fmt"
func maxOne(arr []int , k int){
	left:=0
	zeroCount:=0
	maxValue:=0

for right:=0;right<len(arr);right++{

	if arr[right]==0{
		zeroCount++
	}

	if zeroCount>k{
		if arr[left]==0{
			zeroCount--
		}
    left++

	}


	if maxValue<right-left+1{
		maxValue=right-left+1

	}


}

fmt.Println(maxValue)


}


func main(){

	arr :=[]int{1,1,1,0,0,0,1,1,1,1,0}
	k:=2

	maxOne(arr,k)
}