package main


import "fmt"

func signs(arr []int , k int){

	for i:=0;i<len(arr);i++{

		for j:=1;j<len(arr)-1;j++{

			if arr[i]+arr[j]==k{
				fmt.Println(arr[i],arr[j])
			}
		}
	}


}


func main(){
	arr:=[]int{2,3,5}
	k:=5

	signs(arr,k)
}