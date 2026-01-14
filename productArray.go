package main

import "fmt"

func prductArray(arr []int){

n := len(arr)
res :=make([]int , n)


left:=1

for i:=0;i<n;i++{
	res[i]=left
	left *=arr[i]
}

right:=1

for i:=n-1; i>=0;i--{
	res[i]*=right
	right*=arr[i]
}

fmt.Println(res)
}


func main(){
	arr :=[]int{1,2,3,4}
	prductArray(arr)
}