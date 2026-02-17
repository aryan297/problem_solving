package main

import ("sort"
         "fmt")
func threeSum(arr []int ,target int){

	sort.Ints(arr)

	sum:=0

	result :=[][]int{}

	for i:=0;i<len(arr)-2;i++{

	    if i>0 && arr[i]==arr[i-1]{
	    	continue
	    }

		left:=i+1
		right:=len(arr)-1

		for left <right{

			sum = arr[i]+arr[left] +arr[right]

			if sum==target{
				result=append(result, []int{arr[i],arr[left],arr[right]})
				fmt.Println([]int{arr[i],arr[left],arr[right]})
				left++
			right--
			

			for left <right && arr[left]==arr[left-1]{
				left++
			}
			for left <right && arr[right]==arr[right-1]{
				 right--
			}

			}else if sum < target{
				left++
			} else{
				right--
			}

		}


	}

	fmt.Println(result)


}


func main(){
	arr :=[]int{-1,0,1,2,-1,-4}

	threeSum(arr,0)
}