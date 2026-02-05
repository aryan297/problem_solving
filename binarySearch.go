package main

import "fmt"


func binarySearch(arr []int , target int) int{

	left , right :=0, len(arr)-1

	for left<=right{
		mid:= left+(right-left)/2

		fmt.Println(mid)

		if arr[mid]==target{
			return mid
		} else if arr[mid]<target{
			left=mid+1

		}else{
          right=mid-1

		}

	}

	return -1
}


func main(){
	arr :=[]int{1,4,8,9,10}

	k := binarySearch(arr,8)

	fmt.Println(k , "output")
}