package main

import "fmt"
func zeroOneMatrix(mat [][]int){

	row , col := len(mat) , len(mat[0])

	queue := [][]int{}


	for i:=0;i<row;i++{

		for j:=0;j<col;j++{

		if mat[i][j]==0{
			queue=append(queue, []int{i,j})
		} else{
			 mat[i][j]=-1

		}

	}
}


dict :=[][]int{{0,1},{0,-1},{1,0},{-1,0}}

for len(queue)>0{

	node :=queue[0]
	queue =queue[1:]

	for _ , d := range dict{
		nr := node[0]+d[0]
		nc := node[1]+d[1]

		if nr < 0 || nc < 0 || nr >= len(mat) || nc >= len(mat[0]) || mat[nr][nc]!=-1{
			continue
		}
		  mat[nr][nc]= mat[node[0]][node[1]] +1
		  queue = append(queue, []int{nr, nc})

		}


}

fmt.Println(mat)

}


func main(){

	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	zeroOneMatrix(mat)
}