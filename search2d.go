package main


import "fmt"



func search(arr [][]int,target int) bool{

	row:=len(arr)
	col:=len(arr[0])

	left:=0
	right:=row*col-1

	fmt.Println(right ,left)

	for left<=right{
		mid:=left+(right-left)/2
		row=mid/col 
		col=mid%col
		midVal:=arr[row][col]

		if midVal==target{
			fmt.Println(row,col)
			return  true
		} else if midVal < target{
			left=mid+1

		} else{
			right=mid-1
		}
}

return false


}


func main(){

	arr :=[][]int{{1,  3,  5,  7},
	{10, 11, 16, 20},
	{23, 30, 34, 60}}

	target:=11

	search(arr ,target)

}


