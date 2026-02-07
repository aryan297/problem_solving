package main


import "fmt"

func graphEdge(n int , edges [][]int) int{

	graph := make([][]int , n)


	for i:=0;i<len(edges);i++{
		u , v := edges[i][0], edges[i][1]


	graph[u] = append(graph[u],v)
	graph[v] = append(graph[v],u)
    }

    visited :=make([]bool , n)
    result:=0


    var dfs func(int, *[]int , *int)


    dfs = func(node int , nodes *[]int, edgeCount *int){

    	visited[node]=true

    	*nodes = append(*nodes , node)

    	*edgeCount += len(graph[node])

    	for _ , neigh := range graph[node]{
    	if !visited[neigh]{
    		dfs(neigh,nodes,edgeCount)
    	}
    	}

    }


    for i:=0;i<n; i++{

    	if !visited[i]{

    	nodes:=[]int{}
    	edgeCount :=0


    	dfs(i,&nodes,&edgeCount)


    	k:=len(nodes);

    	actualEdge :=edgeCount/2

    	requiredEdge :=k*(k-1)/2

    	if actualEdge==requiredEdge{
    		result++
    	}
    }

    }

    return result

}


func main(){
	n := 6
	//edges = [[0,1],[0,2],[1,2],[3,4]]

	edges :=[][]int{{0,1},{0,2},{1,2},{3,4}}
	val := graphEdge(n,edges)

	fmt.Println(val)


}