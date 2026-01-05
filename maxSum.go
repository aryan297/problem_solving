package main

import "fmt"


func maxSum(arr []int , k int){



	maxSum:=0
			sum:=0
	for i:=0;i<k;i++{

		sum = sum+arr[i]

		maxSum=sum
	}

	for i:=k;i<len(arr);i++{


		sum =sum + arr[i] - arr[i-k]
       maxSum=max(maxSum,sum)

	}

  fmt.Println(maxSum)

}


func max(a, b int) int {

	if a >b{
		return a
	}

	return b
}



func main(){

	arr :=[]int{2,1,5,1,3,2}
	k:=3
	maxSum(arr,k)
}