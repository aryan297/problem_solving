package main

import "fmt"


func longest(arr []int , k int){
	left:=0
	maxCount:=0
	zeroCount:=0

	for right:=0;right<len(arr) ;right++{

		if arr[right]==0 {

		   zeroCount++

		}
		for zeroCount >k {

			if arr[left]==0{
			zeroCount--

			}
			left++
		}

		if right-left+1>maxCount{
			maxCount=right-left+1
		}

	}

	fmt.Println(maxCount)


}

func main() {
	arr := []int{1,1,1,0,0,0,1,1,1,1,0}
	k:=2

	longest(arr,k)
}