package main

import "fmt"

func dutch(arr []int){

	low :=0
	mid:=0
	high:=len(arr)-1

	for mid <high{
		switch arr[mid] {
		case 0:
			arr[low] ,arr[mid]=arr[mid],arr[low]
			low++
			mid++

		case 1:
			mid++
		case 2:
			arr[mid], arr[high] =arr[high],arr[mid]
			high--

		}
			
		
		
	}

}


func main(){
	arr :=[]int{2, 0, 1, 2, 1, 0, 1, 2, 0, 1}
	dutch(arr)
	fmt.Println(arr)
}