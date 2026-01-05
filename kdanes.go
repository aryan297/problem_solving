package main

import "fmt"


func kdanes(arr []int){

	curSum:=0
	maxSum:=0
	start ,end:=0,0

	temp:=0
	for i:=0;i<len(arr);i++{

		curSum =curSum + arr[i]

		if curSum>maxSum{
			maxSum=curSum
			start=temp
			end=i
		}

		if curSum<0{
			curSum=0
			temp=i+1
		}

	}

	fmt.Println(arr[start:end+1])

}



func main(){
	arr:=[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	kdanes(arr)
}