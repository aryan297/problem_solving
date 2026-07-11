package main


import "fmt"

func bestTime(arr []int){

	minPrice :=arr[0]
	maxProfit :=0

  for i :=1;i<len(arr);i++{

  	if arr[i]<minPrice{
  		minPrice=arr[i]
  	}

  	profit :=arr[i]-minPrice

  	
  	if profit>maxProfit{
  		maxProfit=profit
  	}
}

fmt.Println(maxProfit)

}
func main(){

	arr :=[]int{7, 1, 5, 3, 6, 4}

	bestTime(arr)
}