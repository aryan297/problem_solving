package main

import "fmt"


func moveZero(arr []int){

	left :=0

	for right:=0;right<len(arr);right++{

       if arr[right]!=0{
       	arr[left]=arr[right]
      left++
       }
	}

	for i:=left;i<len(arr);i++{
		arr[i]=0
	}

	fmt.Println(arr)


}



func moveZ(arr []int){

	left:=0

		for right:=0;right<len(arr);right++{


		if arr[right]!=0{

			arr[left],arr[right]=arr[right],arr[left]
			left++
		}


	}
	fmt.Println(arr)

}


func main(){

	arr := []int{1,7,3,0,3,0,1}
	moveZero(arr)
	moveZ(arr)

}