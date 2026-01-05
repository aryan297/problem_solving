package main

import "fmt"



func transpose(arr [][]int){
	row:=len(arr)
	col:=len(arr[0])

	transposed:=make([][]int ,col)

	for i:=range transposed{
		transposed[i]=make([]int ,row)
	}

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			transposed[j][i]=arr[i][j]
		}
	}

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{

		}
	}

	for i:=0;i<row;i++{
		left:=0
		right:=col-1

		for left<=right{
			transposed[i][left],transposed[i][right]=transposed[i][right],transposed[i][left]
			left++
			right--
		}
	}


	fmt.Println(transposed)




}

func main(){

	matrix :=[][]int{{1,2,3},{4,5,6} ,{7, 8, 9}}
	transpose(matrix)


}