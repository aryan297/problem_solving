package main

import "fmt"


func permutation(arr []int){

	var res []int

	var path [][]int

	used := make([]bool,len(arr))

	var dfs func()

	dfs=func(){

		if len(res)==len(arr){

			temp:=append([]int{}, res...)
			path=append(path,temp)

			return 
		}

		for i:=0;i<len(arr);i++{

			if used[i]{
				continue
			}

		  if i>0 && arr[i]==arr[i-1] && !used[i-1]{
		  	continue
		  }

		  used[i]=true
		  res=append(res,arr[i])

		  dfs()
		  res=res[:len(res)-1]
		  used[i]=false
		}

	}

	dfs()

fmt.Println(path)


}


func main(){
	arr :=[]int{1,2,3}
	permutation(arr)

}