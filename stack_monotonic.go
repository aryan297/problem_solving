package main

import "fmt"


func monotonic(arr []int) map[int]int{

	stack := []int{}
	ans := map[int]int{}

	for _, num := range arr{

		for len(stack) >0 && num > stack[len(stack)-1] {

		top:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans[top]=num
	}

	stack=append(stack,num)
}


for len(stack)>0{

	top := stack[len(stack)-1]

	stack =stack[:len(stack)-1]

	ans[top]=-1

}

fmt.Println(ans)

return ans

}

func main(){

	arr := []int{4,2,6,3,8}
	k :=monotonic(arr)

	for _,v := range k{
		fmt.Println(v)
	}
}