package main

import "fmt"



func maxProduct(arr []int){

	n:=len(arr);

	if n==0{
		fmt.Println()
	}

	ans:=arr[0]


	prefix,suffix:=1,1


	for i:=0;i<n;i++{

		if prefix==0{

			prefix=1
		}

		prefix*=arr[i]


		if suffix==0{
			suffix=1
		}

		suffix*=arr[n-i-1]


    ans=max(ans, max(prefix,suffix))

	}


	fmt.Println(ans)
}

func max(a,b int) int{

	if a>b{
		return a
	}

	return b
}


func main(){
	arr :=[]int{-2,3,-4}

	maxProduct(arr)
}