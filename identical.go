package main

import "fmt"

func identicle(arr []int){

frq := make(map[int]int)

pairs :=0

for i:=0;i<len(arr);i++{

	pairs += frq[arr[i]]




	frq[arr[i]]++
}

fmt.Println(pairs)
fmt.Println(frq)

}
func main(){
	arr :=[]int{1, 2, 3, 1, 1, 3}

	identicle(arr)
}