package main

import "fmt"

func topoSort( n int , edges [][]int){

	graph := make(map[int][]int)

	indegree := make([]int,n)

	queue := []int{}

	for _, e := range edges{
		u , v :=e[0],e[1]
		graph[v]=append(graph[v],u)
		indegree[u]++
	}

	for i:=0;i<n;i++{
		if indegree[i]==0{
			queue=append(queue,i)
		}
	}


	for len(queue)>0{
		node := queue[0]
		queue=queue[1:]

		fmt.Println(node)

		for _ , neigh := range graph[node]{
			indegree[neigh]--

			if indegree[neigh]==0{
				queue=append(queue,neigh)
			}
		}
	}


}

func main(){
	edges :=[][]int{{1,0},{2,0},{3,1},{3,2}}
	topoSort(4,edges)
}