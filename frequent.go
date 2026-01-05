package main

import "fmt"


func countFreq(arr []int){

	hashMap:=make(map[int]int)

	for i:=0;i<len(arr);i++{
		hashMap[arr[i]]++
	}

	maxFeq:=0
	result:=arr[0]


	for num , freq := range hashMap{
		if freq> maxFeq || ( freq == maxFeq && num<result) {
			maxFeq=freq
			result=num
		}
	    

	}
 
 fmt.Println(maxFeq , result)


}


func main(){
	arr:=[]int{1, 3, 2, 3, 3, 2, 3, 1}

	countFreq(arr)
}