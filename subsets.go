package  main

import "fmt"


func subset(arr []int) [][]int{

	var res [][]int
	var path []int
	var dfs func(int)
	dfs = func(start int){
		temp :=append([]int{}, path...)
		res = append(res, temp)

		for i:=start;i<len(arr);i++{

			path =append(path , arr[i])
			dfs(i+1)
			path =path[:len(path)-1]
			
		}

	}

	dfs(0)

	fmt.Println(res)
	return res

}

func main(){


	arr :=[]int{1,2,3}
	k:=subset(arr)

	fmt.Println(k)
}