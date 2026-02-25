package main

import "fmt"


func maxProfit(arr []int){

	max:=0

	minPrice:=arr[0]

	for i:=1;i<len(arr);i++{

		if arr[i]<minPrice{
			minPrice=arr[i]
		}

		profit:=arr[i]-minPrice

		if profit>max{
			max=profit
		}

	}

	fmt.Println(max)


}



func main() {
		arr := []int{7, 1, 5, 3, 6, 4}

		maxProfit(arr)
}