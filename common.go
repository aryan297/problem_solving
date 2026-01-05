package main

import "fmt"

func longestCommon(arr []int){

	hashMap:=make(map[int]bool);
	longest:=0

	for i:=0;i<len(arr);i++{

		hashMap[arr[i]]=true
	}

	fmt.Println(hashMap)

	for num := range hashMap{

		if !hashMap[num-1]{
			length:=1
			cur:=num


		for hashMap[cur+1]{
			cur++;
			length++;
		}

		if length>longest{
			longest=length
		}

		}
	}

	fmt.Println(longest)
}


func main(){
	arr :=[]int{100, 4, 200, 1, 3, 2}
	longestCommon(arr)
}