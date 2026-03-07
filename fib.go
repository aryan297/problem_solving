package main

import "fmt"

func fib(n int) int{

	if n<=1{
		return n
	}


	prev1:=1
	prev2:=0

	for i :=2;i<n;i++{
		cur:=prev1+prev2

		prev2=prev1

		prev1=cur
	}
return prev1

}


func main(){

	k:=fib(3)

	fmt.Println(k)

}