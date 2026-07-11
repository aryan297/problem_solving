package main


import "fmt"


func abs(a  int) int{

	 if a <0{
   return -a

	 }

	 return a
}




func dispair(arr []int){

	for i:=0;i<len(arr);i++{

		index:=abs(arr[i])-1


		if arr[index]>0{
			arr[index]= -arr[index]
		}


	}
  fmt.Println(arr)

}


func main(){

	arr :=[]int{4, 3, 2, 7, 8, 2, 3, 1}
	dispair(arr)
}