package main


import "fmt"


func minSum(arr []int , target int){

    sum :=0

	left:=0

	minlen:=len(arr)+1

	for right:=0;right<len(arr);right++{

		sum +=arr[right]

		for sum>=target{

			if right-left+1<minlen{
				minlen=right-left+1
			}

			sum-=arr[left]
			left++

		}


	}

	fmt.Println(minlen)


}

func main(){
	arr :=[]int{2,3,1,2,4,3}
	minSum(arr,7)
}