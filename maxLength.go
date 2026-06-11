package main

import "fmt"

func maxLength(arr []int){

	seen := map[int]int{0:-1}

	sum, best :=0,0

    for i ,v := range arr{

    	if v==0{
    		sum--
    	} else{
    		sum++
    	}

    	if idx, ok :=seen[sum];ok{
    		if i-idx>best{
    			best=i-idx
    		}
    	} else{
    		seen[sum]=i
    	}

    }


fmt.Println(best)


}


func subarray(arr []int){

	best := arr[0]
	cur := arr[0]

	for _, n := range arr[1:]{
		if cur+n >n{
			cur=cur+n
		} else{
			cur=n
		}

		if cur>best{
			best=cur
		}

	}

	fmt.Println(best)


}


func main(){
	arr := []int{0, 0, 1, 0, 0, 0, 1, 1}

	arr2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxLength(arr)
	subarray(arr2)
}