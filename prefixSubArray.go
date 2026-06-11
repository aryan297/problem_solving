package main


import "fmt"


func subArray(arr []int , k int){

	sumMap:= map[int]int{0:-1}
	sum , best :=0,0
	for i, v := range arr{
		sum +=v
		if idx , ok :=sumMap[sum-k];ok{

			if i-idx > best{
				best=i-idx
			}
		}

		if _,ok :=sumMap[sum]; !ok{

			sumMap[sum]=i
		}

	}

fmt.Println(best)


}


func main(){ 
	arr :=[]int{1, -1, 5, -2, 3}
	k:=3

	subArray(arr,k)
}