package main

import "fmt"
func consSum(arr []int , k int){
	left,sum,maxlen:=0,0,0

	for right:=0;right<len(arr);right++ {

		sum +=arr[right]


		for sum > k && left <=right{

			sum-=arr[left]
			left++
		}

		if sum==k{

			if right-left+1>maxlen{
				maxlen=right-left+1
			}
		}



	}

	fmt.Println(maxlen)


}


func main(){
	arr:=[]int{2,3,5,1,9}

	consSum(arr,10)


}