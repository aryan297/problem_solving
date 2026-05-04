package main

import "fmt"

func solveNqueen(n int) {

	var ans [][]string

	board := make([][]byte , n)

	for i:=0;i<n;i++{
		board[i]=make([]byte,n)
		for j:=0;j<n;j++{
			board[i][j]='.'
		}
	}

	cols :=make(map[int]bool)
	diag1:=make(map[int]bool)
	diag2 :=make(map[int]bool)

	var dfs func(int)

	dfs= func(row int){

		if row==n{
			temp:=make([]string,n)
			for i:=0;i<n;i++{

				temp[i]=string(board[i])
			}
			ans=append(ans,temp)
			return
		}

		for col:=0;col<n;col++{

			if cols[col] || diag1[row+col] || diag2[row-col]{
				continue
			}

			board[row][col]='Q'
			cols[col]=true
			diag1[row+col]=true
			diag2[row+col]=true

			dfs(row+1)

			board[row][col]='.'
			cols[col]=false
			diag1[row+col]=false
			diag2[row+col]=false


		}


	}

	dfs(0)


	fmt.Println(ans)



}

func main() {
solveNqueen(4)
}