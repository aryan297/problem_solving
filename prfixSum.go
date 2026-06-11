package main

import "fmt"

func subArray(arr []int , k int){

	countMap := make(map[int]int)

	countMap[0]=1

	res , sum :=0,0

	for _ , v:= range arr{
		sum+=v

		res+=countMap[sum-k]

		countMap[sum]++
	}

	fmt.Println(countMap)
	fmt.Println(res)


}

func main() {
	arr := []int{1, 1, 1}
	k:=2

	subArray(arr,k)
}