package main

import "fmt"


func indexSum(arr []int , target int){

	sum:=0;
	left:=0

	for right:=0;right<len(arr);right++{

		sum +=arr[right]

		for sum > target  && left <=right{

			sum -=arr[left]
			left++


			if sum==target{
				fmt.Println(left+1,right+1)
				return
			}
		}
	}



}


func main(){

	arr :=[]int{1,2,3,7,5}

	target:=12

	indexSum(arr,target)
}