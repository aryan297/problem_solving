package main

import "fmt"



func once(arr []int){

	res:=0

	for i:=0;i<len(arr);i++{
		res^=arr[i]
	}

	fmt.Println(res)
}

func main(){

	arr :=[]int{1,2,5,6}
	once(arr)
}