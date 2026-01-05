package main

import "fmt"

func sexMatrix(matrix [][]int){
	m:=len(matrix)
	n:=len(matrix[0])
	firstRowZero := false
	firstColZero := false

	for i:=0;i<m;i++{
		if matrix[i][0]==0{
			firstColZero=true

		}

	}

	for j:=0;j<n;j++{
		if matrix[0][j]==0{
			firstRowZero=true
		}
	}


	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			if matrix[i][j]==0{
				matrix[i][0]=0
				matrix[0][j]=0
			}

		}
	}
	// Step 3: Set zeroes based on marks
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}


	if firstColZero{
		for i:=0;i<m;i++{
			matrix[i][0]=0
		}
	}

	if firstRowZero{
		for j:=0;j<n;j++{
			matrix[0][j]=0
		}
	}



}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}


func main(){

	matrix :=[][]int{{1,1,1} , 
	                 {1,0,1},
		             {1,1,1}}

sexMatrix(matrix)

printMatrix(matrix)


}