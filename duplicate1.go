package main

import ("fmt" 
        "math")


func findDuplicate(arr []int){

	duplicate:=[]int{}

for i:=0;i<len(arr);i++{

	val:=int(math.Abs(float64(arr[i])))

	if arr[val-1]<0{
		duplicate=append(duplicate,val)
	} else{
		arr[val-1]=-arr[val-1]
	}
}
 
fmt.Println(duplicate)
}

func main(){

	nums := []int{3, 1, 2, 5, 3 , 2}
   findDuplicate(nums)
}