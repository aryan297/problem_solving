package main

import "fmt"


func loop(arr [][]int , target int) int{

	n:=len(arr)
	m:=len(arr[0])

	row , col := 0 , m-1

	for row < n && col >=0{

		if arr[row][col]==target{
			fmt.Println(row, col)
			return 1
		} else if arr[row][col]>target{
			col=col-1
		} else{
			row=row+1
		}

	}
   return -1

}



func main(){

	arr:=[][]int{
		      {1,2,3,4},
              {5,6,7,8},
              {9,10,11,12},
              }

    k:=loop(arr,12)

   fmt.Println(k)

}