package main

import "fmt"

func productArray(arr []int){
    n:=len(arr)
	ans := make([]int, n)

	ans[0]=1
	for i:=1;i<len(arr);i++{
      ans[i] = ans[i-1] * arr[i-1]

	}
	right:=1
for i:=n-1;i>=0;i--{
	ans[i]*=right;
	right*=arr[i]

}

	fmt.Println(ans)
}

func main(){
	arr :=[]int{1,2,3,4} // {2*3*4, 1*3*4, 1*2*4, 1*2*3}
	productArray(arr)

}