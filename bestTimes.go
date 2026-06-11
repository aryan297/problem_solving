package main

import "fmt"

func bestTime(arr []int){

	minPrice :=arr[0]
	best:=0

	for i:=0;i<len(arr);i++{

		if arr[i]< minPrice{
			minPrice=arr[i]
		} else if arr[i]-minPrice > best{
			best=arr[i]-minPrice
		}
	}

	fmt.Println(best)


}

func main(){
	arr :=[]int{7, 1, 5, 3, 6, 4}
	bestTime(arr)
}