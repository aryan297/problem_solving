package main

import "fmt"


func transpose(arr [][]int){

	for i:=0;i<len(arr);i++{
		for j:=i+1; j<len(arr);j++{
			arr[i][j],arr[j][i]= arr[j][i],arr[i][j]
		}
	}

	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)/2 ;j++{
			arr[i][j],arr[i][len(arr)-1-j]=arr[i][len(arr)-1-j],arr[i][j]
		}
	}

	//fmt.Println(arr)

}


func printArray(arr [][]int){

	for i:=0;i<len(arr);i++{

		fmt.Println(arr[i])
	
	}
}

func main(){
	arr :=[][]int{ 
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	    }

	transpose(arr)
	printArray(arr)
}