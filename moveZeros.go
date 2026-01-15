package main

import "fmt"


func moveZero(arr []int){
	j:=0
	for i:=0;i<len(arr);i++{
		if arr[i]!=0{
			arr[i],arr[j]=arr[j],arr[i]
			j++
		}
	}
	fmt.Println(arr)
}

func main(){
	arr := []int{2,3,0,0,2}

	moveZero(arr)
}