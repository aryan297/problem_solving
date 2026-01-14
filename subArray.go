package main

import "fmt"


func subArraySum(arr []int , k int){

	count:=0
	sum:=0

	hashMap :=map[int]int{0:1}

	for i:=0;i<len(arr);i++{

		sum +=arr[i];

		if hashMap[sum-k]>0{
			fmt.Println(sum, k, "k")
			fmt.Println(hashMap[sum-k])

			count +=hashMap[sum-k]
		}


      hashMap[sum]++


	}

fmt.Println(count , "count")



}



func main(){
	arr :=[]int{1 ,2,3}
	k:=3
subArraySum(arr,k)
}