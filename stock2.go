package main

import "fmt"


func maxProfit(arr []int){

	profit:=0

	for i:=1;i<len(arr);i++{

		if arr[i]>arr[i-1]{

			profit+=arr[i]-arr[i-1]
		}
	}


	fmt.Println(profit)


}
func main(){
		arr := []int{7, 1, 5, 3, 6, 4}
		maxProfit(arr)
}