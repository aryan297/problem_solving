package main


import "fmt"



func arrData(arr []int){

	start,mid,high:=0,0,len(arr)-1

	for mid<=high {
		switch  arr[mid] {
	
		
	case 0:
		arr[start],arr[mid]=arr[mid],arr[start]
		start++;
		mid++
	case 1:
		mid++;

	case 2:
		arr[mid],arr[high]=arr[high],arr[mid]
		high--
	}
	}


fmt.Println(arr)


}


func main(){

	arr :=[]int{1,2,0,1,2}
	arrData(arr)
}