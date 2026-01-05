package main

import "fmt"

func sorted(arr []int) bool{

	for i:=0;i<len(arr);i++{
		k:=i+1

		if k!=arr[i]{

			return false
		}

	}

	return true

}


func main(){

	arr :=[]int{1,2,3,4,5}

	fmt.Println(sorted(arr))

}