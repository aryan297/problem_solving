package main

import "fmt"

func union(arr1 , arr2 []int){

	set :=make(map[int]bool)

	for _ , v:=range arr1{
		set[v]=true
	}

    for _ , v:=range arr2{
    	set[v]=true
    }

    result:=[]int{}

    for  val :=range set{
    	result=append(result,val)
    }

    fmt.Println(result)


}


func main(){
	arr1 :=[]int{1, 2, 3, 4, 5}
	arr2 :=[]int{2, 3, 4, 4, 5}

	union(arr1,arr2)
}