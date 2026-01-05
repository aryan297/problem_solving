package main

import ("fmt"
	    "sort")




func threeSum(arr []int){
    
    sort.Ints(arr)

    fmt.Println(arr)

    res:=[][]int{}

    for i:=0 ; i<len(arr)-2 ; i++{

    	if i >0 && arr[i]==arr[i-1]{
    		continue
    	}

    	l,r :=i+1, len(arr)-1

    	for l<r{
    		sum :=arr[i]+arr[l]+arr[r]

    		if sum==0{
    			res=append(res, []int{arr[i],arr[l],arr[r]})

    			for l<r && arr[l]==arr[l+1]{
    				l++
    			}

    			for l<r && arr[r]==arr[r-1]{
    				r--
    			}

    			l++
    			r--
    		} else if sum <0{
    			l++
    		} else{
    			r--

    		}


    	}



    }

    fmt.Println(res)




}

func main(){
	arr :=[]int{-1, 0, 1, 2, -1, -4}
	threeSum(arr)
}