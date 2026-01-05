package main

import "fmt"



func spiralGo(arr [][]int){

	top , bottom :=0, len(arr)-1
	left , right :=0, len(arr[0])-1
	res:=[]int{}


	for left <= right && top <=bottom{

		for i:=left;i<=right;i++{
		  res=append(res,arr[top][i])
		}
		top++;

		for i:=top;i<=bottom;i++{
			res=append(res,arr[i][right])
		}
		right--

		if top<=bottom{
			for i:=right; i>=left;i--{
				res=append(res,arr[bottom][i])
			}
			bottom--
		}

		if left<=right{
			for i:=bottom;i>=top;i--{
				res=append(res,arr[i][left])
			}
			left++
		}

	}

	fmt.Println(res)

}



func main(){

	arr :=[][]int{{1,2,3},
                  {4,5,6},
                   {7,8,9}}

     spiralGo(arr)

}