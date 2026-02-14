package main

import "fmt"

func leftPart(arr []int ,target int) int {
	left, right := 0, len(arr)-1
	ans:=-1

	for left<=right{

		mid:=left+(right-left)/2

		if arr[mid]==target{
			ans=mid
			right=mid-1

		} else if arr[mid] < target{
			right=mid-1

		} else {
			left=mid+1
		}
	}

	return ans
	
}


func rightPart(arr []int ,target int) int {
	left, right := 0, len(arr)-1
	ans :=-1

	for left<=right{

		mid:=left+(right-left)/2


		if arr[mid]==target{
			ans=mid
			left=mid+1
		} else if arr[mid] < target{
			left=mid+1

		} else {
			right=mid-1
		}
	}
	return ans
}



func main(){
	arr :=[]int{1,3,4,4,5,6}


	l:=leftPart(arr,4)
	r := rightPart(arr,4)

	fmt.Println(l,r)
}