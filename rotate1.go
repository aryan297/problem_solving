package main

import "fmt"

func rotate(arr []int , k int){
	n:=len(arr)-1
	k=k%n

	reverse(arr,0,n-1)
	reverse(arr,0,k-1)
	reverse(arr,k,n-1)
   
   fmt.Println(arr)

}


func reverse(arr []int , left , right int){

	for left <=right{

		arr[left] , arr[right] =arr[right] , arr[left]
		left++
		right--
	}

}





func main(){
	arr :=[]int{1,2,3,4,5,6,7}
	k:=3
	rotate(arr,k)


}