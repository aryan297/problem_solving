package main

import "fmt"




func leftRotate(arr []int, k int) []int{
     
	if len(arr)==0{
		return arr
	}

	K :=k% len(arr)

	reverse(arr , 0 , K-1)

	reverse(arr ,K , len(arr)-1)

	reverse(arr , 0 , len(arr)-1)


   return arr;

}

func reverse(arr []int ,left , right int){

	for left <right{
		arr[left],arr[right]=arr[right],arr[left]
		left++
		right--
	}

	fmt.Println(arr)
}



func main(){
	arr:=[]int{1,2,3,4,5}

	fmt.Println(leftRotate(arr,1))
}


