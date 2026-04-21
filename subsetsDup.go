package main

import ("fmt"
	    "sort")

func subset(arr []int) {

	sort.Ints(arr)

	res :=[][]int{}
	path :=[]int{}

	var dfs func(int)

	dfs = func(start int) {

		temp:=append([]int{} , path...)
		res =append(res,temp)

		for i:=start;i<len(arr);i++{

			if i > start && arr[i]==arr[i-1]{
				continue
			}

			path=append(path,arr[i])
			dfs(i+1)
			path=path[:len(path)-1]

		}

	}
	dfs(0)

	fmt.Println(res)


}

func main(){

	arr := []int{1,2,2}
	subset(arr)
}