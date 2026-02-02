package main

import "fmt"


func dfs(node int , graph map[int][]int , visited map[int]bool){

	if visited[node]{
		return
	}

	visited[node]=true

	fmt.Println(node)

	for _,neighbour := range graph[node]{
		dfs(neighbour,graph,visited)
	}
}


func bfs( start int , graph map[int][]int){
	visited := make(map[int]bool)

	queue:=[]int{start}

	visited[start]=true

	for len(queue)>0{
		node:=queue[0]
		queue=queue[1:]

		fmt.Println(node , "bfs")

		for _, neighbour := range graph[node]{

			if !visited[neighbour]{
				visited[neighbour]=true

				queue=append(queue,neighbour)
			}
		}


	}
}



func main(){


	graph :=make(map[int][]int)

    graph[0] = []int{1, 2}
	graph[1] = []int{2}
	graph[2] = []int{0, 3}
	graph[3] = []int{3}

	fmt.Println(graph)

	visited := make(map[int]bool)
		fmt.Println("DFS traversal starting from node 2:")
	dfs(2,graph,visited)
		fmt.Println()


    bfs(0,graph)






}