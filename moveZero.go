package main

import "fmt"

func move(arr []int){
	j:=0

	for i:=0;i<len(arr);i++{

		if arr[i]!=0 {

		if i!=j{
			arr[i],arr[j]=arr[j],arr[i]
		}
		j++
	}

	}

	fmt.Println(arr)
}

func main(){

	arr :=[]int{1, 0, 2, 3, 0, 4, 0, 1}
	move(arr)
}