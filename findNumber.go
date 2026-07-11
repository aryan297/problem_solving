package main

import ("fmt"
        "strconv")


func findNum(arr []int){

	count :=0

	for i:=0;i<len(arr);i++{

		if len(strconv.Itoa(arr[i]))%2==0{
			 count++
		}


	}

	fmt.Println(count)
}


func wealth(arr [][]int){

	maxWealth:=0

	for i:=0;i<len(arr);i++{
		wealth :=0
		for j:=0;j<len(arr[i]);j++{
			wealth +=arr[i][j]

		}

		if wealth>maxWealth{
			maxWealth=wealth
		}
	}

  fmt.Println(maxWealth)
}

func main(){
	arr := []int{12,345,2,6,7896}
	findNum(arr)
	acc := [][]int{{1, 5},{7,3},{3,5}}
	wealth(acc)

}