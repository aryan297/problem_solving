package main

import "fmt"

func numsLand(grid [][]byte) int{

	if len(grid)==0{
		return 0
	}
  count :=0
	row , col := len(grid) , len(grid[0])

	dict :=[][]int{{1,0},{-1,0},{0,1},{0,-1}}

	for i:=0;i<row;i++{
		for j:=0; j<col;j++{

			if grid[i][j]=='1'{
				count++;
				queue :=[][]int{{i,j}}
  
				grid[i][j]='0'

				for len(queue)>0{

					r ,c := queue[0][0],queue[0][1]
				   queue = queue[1:]

				   for _ , d := range dict{
				   	nr := r+d[0]
				   	nc :=  c+d[1]

				   	if nr>=0 && nc >=0 && nr <row && nc <col && grid[nr][nc]=='1'{

				   		grid[nr][nc]='0'
				   		queue=append(queue,[]int{nr,nc})

				   	}

				   }
				}

			}

		}
	}

	return count

}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

   k:=numsLand(grid)
   fmt.Println(k)

}