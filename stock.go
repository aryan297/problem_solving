package main

import "fmt"

func stockArray(arr []int){

	price:=arr[0]
	maxProfit:=0

	for i:=1;i<len(arr);i++{

		if arr[i]<price{
		  price=arr[i]
		}else{
			profit:=arr[i]-price
			if profit>maxProfit{
				maxProfit=profit
			}
		}


	}

	fmt.Println(maxProfit)


}



func main(){
	arr :=[]int{7,1,5,3,6,4}

	stockArray(arr)
}