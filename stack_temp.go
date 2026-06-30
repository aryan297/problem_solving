package main

import "fmt"


func stack_data(arr []int ) {

	stack := []int{}

	ans := make([]int,len(arr))

	for i:=0;i<len(arr);i++{

		for len(stack)>0 && arr[i]>arr[stack[len(stack)-1]]{

			top :=stack[len(stack)-1]
			stack=stack[:len(stack)-1]

			ans[top]=i-top
		}

		stack=append(stack,i)
	}

fmt.Println(ans)


}


func main(){

	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}

	stack_data(arr)
}