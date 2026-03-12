package main

import "fmt"


func twoSum(arr []int , target int){

	hashMap := make(map[int]int)

	for i:=0;i<len(arr);i++{


		if exist,ok := hashMap[target-arr[i]];ok{
             
           fmt.Println(exist,i)
		}

          hashMap[arr[i]]=i

	}



}


func main(){
	arr := []int{1,3,4,5}

	twoSum(arr,9)
}