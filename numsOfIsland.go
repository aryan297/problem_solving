package main


import "fmt"

func numsIsland(grid [][]byte)int{

	row , col := len(grid), len(grid[0])
	count:=0

	var dfs func(int, int)

	dfs = func( r int, c int){

		if r <0 || c <0 || r >= row || c >= col || grid[r][c]=='0'{
			return
		}

		grid[r][c]='0'

		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r ,c+1)
		dfs(r ,c-1)

	}
 
	for i:=0;i<row;i++{
		for j :=0;j<col;j++{

			if grid[i][j]=='1'{
				count++
				dfs(i,j)
			}
		}
	}
return count


}


func main(){
	 grid := [][]byte{
	 	             {'1','1','1','1','0'},
	                 {'1','1','0','1','0'},
	                 {'1','1','0','0','0'},
	                 {'0','0','0','0','0'},
	                 }

	n :=numsIsland(grid)

	fmt.Println(n)
}