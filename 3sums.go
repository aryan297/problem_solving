package main

import ("fmt"
 "sort")


func threeSum(arr []int , target int){

	sort.Ints(arr)

	res :=[][]int{}

	sum:=0

	for i:=0;i<len(arr)-2;i++{

		if i>0 && arr[i]==arr[i-1]{
			continue
		}

		left :=i+1
		right := len(arr)-1


		for left <right{

			sum = arr[i]+arr[left] +arr[right]

		
			if sum==target{

				res=append(res,[]int{arr[i],arr[right],arr[left]})


				for left <right && arr[left]==arr[left+1]{
					left++
				}

				for left <right && arr[right]==arr[right-1]{
				   right--
				}
				left++
				right--
			} else if sum <target{
				left++
			} else{

				right--
			}

		}


	}


	fmt.Println(res)
}

func main(){

	arr :=[]int{-1,0,1,2,-1,-4}

	threeSum(arr,0)
}