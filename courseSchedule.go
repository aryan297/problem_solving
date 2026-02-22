package main


import "fmt"

func courses(numCourses int , preq [][]int ){

	graph :=make(map[int][]int)

	indegree:= make([]int ,numCourses)

	queue := []int{}

	for _ , e := range preq{

	u ,v := e[0],e[1]

	graph[v]=append(graph[v],u)
	indegree[u]++
}

for i:=0;i<numCourses;i++{
	if indegree[i]==0{
		queue=append(queue,i)
	}
}

order :=[]int{}


for len(queue)>0{
	node := queue[0]

	queue = queue[1:]

	order=append(order,node)

	for _ , neigh := range graph[node]{

		indegree[neigh]--

		if indegree[neigh]==0{

			queue=append(queue,neigh)
		}
	}

}


fmt.Println(order)


}


func main() {
	preq := [][]int{{1,0},{2,0},{3,1},{3,2}}

	numCourses := 4
	courses(numCourses,preq)
}