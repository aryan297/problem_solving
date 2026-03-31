package main

import "fmt"


func mostWater(arr []int){
	left, right := 0, len(arr)-1

	maxArea:=0

	for left <right{

		h:=min(arr[left],arr[right])

		w:=right-left
		area:=h*w

		if area>maxArea{
			maxArea=area
		}

		if arr[left]<arr[right] {
			left++
		} else{
			right--
		}
	}

	fmt.Println(maxArea)

}


func min(a, b int) int {

	if a<b{
		return a
	}
	return b

}



func main(){
	arr :=[]int{1,8,6,2,5,4,8,3,7}
	mostWater(arr)
}