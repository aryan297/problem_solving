package main

import "fmt"


func mainArray(arr []int , target int) int{
	left, right:=0, len(arr)-1

	ans:=len(arr)


	for left<=right{
		mid:=left+(right-left)/2

		if arr[mid]==target{
			return mid
		} else if arr[mid]> target{
			ans=mid
			right=mid-1
		} else{
			left=mid + 1
		}
	}
return ans

}


func main(){ß

	arr:=[]int{1, 3, 7, 6} 
	target:=5

	fmt.Println(mainArray(arr, target))



}