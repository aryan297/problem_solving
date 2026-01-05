package main

import ("fmt" 
        "math")


func findMissing(arr []int) (duplicate , missing int){

	for i:=0;i<len(arr);i++{
		val :=int(math.Abs(float64(arr[i])))

		if arr[val-1]<0{
			duplicate=val
			fmt.Println(duplicate,"ss")
		} else {
			arr[val-1]=-arr[val-1]
			fmt.Println(arr[val-1])
		}


	}

	for i:=0;i<len(arr);i++{
		if arr[i]>0{
		missing=i+1

		break
}
	}
   
   return

}


func main(){
	nums := []int{3, -1, 2, 5, 3}
duplicate , missing :=findMissing(nums)
fmt.Println(duplicate ,"dup" , missing, "miss")

}

