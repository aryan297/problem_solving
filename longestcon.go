package main

import ("fmt"
        "sort")


func longestSeq(arr []int){

	sort.Slice(arr, func(i , j int) bool{
		return arr[i]<arr[j]
	})

  
    longest,curr:=1,1

    for i:=1;i<len(arr);i++{

    	if arr[i]==arr[i-1]+1{
    		curr++
    	} else if arr[i]!=arr[i-1]{

    	     if curr>longest{
    	     	longest=curr
    	     }
    	     curr=1
       }

    }
      if curr>longest{
    	     	longest=curr
    	  }

 fmt.Println(longest)
}


func main(){
	arr:=[]int{101,1,3,2,4,102,103,1}
	longestSeq(arr)
}