package main

import "fmt"

func isBitSet(n ,i int) bool {
	mask := 1<<i

	fmt.Printf("binary %b",mask) 

	 return (n&mask)!=0


}

func missingNum(arr []int){

	res := len(arr)

	for i:=0;i<len(arr);i++{
		res ^=i
		res ^=arr[i]
	}

	fmt.Println(res)

}

func main(){

	fmt.Println(isBitSet(13,3))

	arr := []int{0,1,3,4}

	missingNum(arr)



}