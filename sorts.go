package main

import "fmt"

func sort(arr []int){
	left , mid ,right :=0,0, len(arr)-1



	for mid <= right {

		switch arr[mid] {

		case 0:
			arr[left],arr[mid]=arr[mid],arr[left]
			left++
			mid++

		case 1:
			mid++

		case 2:
			arr[mid],arr[right]=arr[right],arr[mid]
			right--
		
		}
	}

	fmt.Println(arr)

}


func main(){

	arr := []int{2, 0, 2, 1, 1, 0}
	sort(arr)
}