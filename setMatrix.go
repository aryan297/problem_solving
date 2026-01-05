package main

import "fmt"


func matrixD(mat [][]int){
	row:=len(mat)
	col:=len(mat[0])
	firstRowZero:=false
	firstColZero:=false


	for j:=0;j<col;j++{
  	if mat[0][j]==0{
  		firstRowZero=true
  		break
  	}
  }
  
  for i:=0;i<row;i++{
  	if mat[i][0]==0{
  		firstColZero=true
  		break;
  	}
  }


  for i:=1;i<row;i++{
  	for j:=1;j<col;j++{
  		if mat[i][j]==0{
  			mat[i][0]=0
  			mat[0][j]=0
  		}
  	}
  }

  for i:=1;i<row;i++{
  	for j:=1;j<col;j++{
  		if mat[i][0]==0 || mat[0][j]==0{
  			mat[i][j]=0
  		}
  	}
  }
  
      // Zero first row if needed
    if firstRowZero {
        for j := 0; j < col; j++ {
            mat[0][j] = 0
        }
    }

    // Zero first column if needed
    if firstColZero {
        for i := 0; i < row; i++ {
            mat[i][0] = 0
        }
    }

  fmt.Println(mat)



}


func main(){
	 matrix := [][]int{
        {1, 1, 1},
        {1, 0, 1},
        {1, 1, 1},
    }
   matrixD(matrix)


}