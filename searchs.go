package main

import "fmt"


func search(arr []int,k int){

	left,right :=0,len(arr)-1


	for left <= right{

		mid :=left+(right-left)/2

		if arr[mid]==k{
			fmt.Println("found" , mid)
			break;
		}else if arr[mid]>k{

			left=left+1

		} else{
			right=right-1
		}

	}

	return 


}


func main(){

	arr := []int{4, 5, 6, 7, 0, 1, 2}

	k:=0

	search(arr,k)


}