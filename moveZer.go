package main

import "fmt"


func moveZero(arr []int){

left :=0

for right :=0;right<len(arr);right++{

	if arr[right]!=0{
		arr[left],arr[right]=arr[right],arr[left]
		left++
	}


}

fmt.Println(arr)
}


func main(){

	arr :=[]int{1,2,0,1,0,2,0}
	moveZero(arr)
}