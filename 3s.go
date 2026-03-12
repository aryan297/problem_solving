package main

import ("sort"
	    "fmt")

func threesum(arr []int , target int){


	sort.Ints(arr)
	arrs:=[][]int{}
 
	for i:=0;i<len(arr)-2;i++{

		if i>0 && arr[i]==arr[i-1]{
			continue
		}

		left:=i+1
		right:=len(arr)-1

		for left<right{
			   sum:=arr[i]+arr[right]+arr[left]

			if sum==target{

				arrs=append(arrs,[]int{arr[i],arr[left],arr[right]})
              
				left++
				right--

				for left<right && arr[left]==arr[left-1]{
					left++
				}


				for left<right && arr[right]==arr[right+1]{
					right--
				}


			} else if sum < target{
				left++
			} else{
				right--
			}


		}

	}

	fmt.Println(arrs)
}

func main(){
	arr := []int{-1,0,1,2,-1,-4}

	threesum(arr,0)
}