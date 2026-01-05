package main

import "fmt"

func duplicate(arr []int) int{

	if len(arr)==0{
		return -1
	}

	j:=0

	for i:=0;i<len(arr);i++{

		if arr[i]!=arr[j]{
			j++;
			arr[j]=arr[i]
		}
	}



return j+1


}

func main(){

	arr :=[]int{1,1,2,2,2,3,3}

   k:=duplicate(arr)
   fmt.Println(k)
    fmt.Println(arr[:k])
   

}