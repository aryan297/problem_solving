package main

import "fmt"


func maxSub(arr []int , k int){
     sum:=0
	for i:=0;i<k;i++{

		sum=sum+arr[i]
	}

	cursum:=sum

	for i:=k;i<len(arr);i++{
    
    sum=sum+arr[i]-arr[i-k]

    cursum=max(cursum,sum)

	}

fmt.Println(cursum)


}


func max(a , b int)int{

	if a>b{
		return a
	}

	return b
}


func main(){
	nums:=[]int{2, 1, 5, 1, 3, 2}
	k:=3

	maxSub(nums,k)
}