package main

import "fmt"

func sorted(arr []int){
left :=0
right :=len(arr)-1
pos:=len(arr)-1
res :=make([]int,len(arr))

for left <right{

	if arr[left] * arr[left] > arr[right]*arr[right] {
		res[pos]=arr[left] *arr[left]
		left++

	} else{

		res[pos]=arr[right]*arr[right]
		right--
	}

	pos--

}

fmt.Println(res)

}

func main(){
	arr :=[]int{4,-1,0,3,10}
	sorted(arr)
}