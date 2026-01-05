package main 

import "fmt"
func majority(arr []int){
	candidate :=arr[0]
	count:=1

	for i:=0;i<len(arr);i++{

		if arr[i]==candidate{
			count++;
		} else{
			count--

			if count==0{
				candidate=arr[i]
				count=1
			}
		}

	}

	fmt.Println(candidate)
}


func main(){
	arr :=[]int{4, 4, 2, 4, 3, 4, 4, 3, 2, 4}

	majority(arr)
}