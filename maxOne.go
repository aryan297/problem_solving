package main

import "fmt"

func ones(arr []int){

	count:=0

	for i:=0;i<len(arr);i++{

		if arr[i]!=1{
          count=0

		} else{
			count=count+1
		}
	}

  fmt.Println(count)

}



func main(){

	arr :=[]int{1, 1, 0, 1, 1,1}
	ones(arr)


}


