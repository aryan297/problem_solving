package main

import "fmt"

func merge(arr1 ,arr2 []int) []int{
	i,j:=0,0

	merged:=[]int{}

	for i<len(arr1) && j<len(arr2) {

		if arr1[i] <= arr2[j]{
			merged=append(merged,arr1[i])
			i++
		} else {
			merged=append(merged,arr2[j])
			j++
		}

	}

	for i<len(arr1) {
		merged=append(merged,arr1[i])
		i++
	}

	for j<len(arr2){
		merged=append(merged,arr2[j])
		j++
	}

	return merged

}


func mergeSort(arr1 []int) []int{

	if len(arr1)<=1{
		return arr1
	}

	mid:=len(arr1)/2

	left:= mergeSort(arr1[:mid])
	right:=mergeSort(arr1[mid:])

	return merge(left,right)

}



func main(){

	arr2:=[]int{7,9,37,3,4}
 fmt.Println(mergeSort(arr2))
}