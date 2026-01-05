package main


import ("fmt"
	    "sort")

func merge(intervals [][2]int){

	sort.Slice(intervals , func(i , j int) bool{
         if intervals[i][0]==intervals[j][0]{

         	return intervals[i][1] < intervals[j][1]
         }

         return intervals[i][0] < intervals[j][0]

	})

	//fmt.Println(intervals)

	merged :=make([][2]int , 0 , len(intervals))
	current :=intervals[0];

	for _ , k :=range intervals[1:]{
		//fmt.Println(k)
		if k[0] <= current[1]{
			if k[1]>current[1]{
			current[1]=k[1]

			}
		} else{
			merged=append(merged,current)
			current=k
		}

	}

	merged=append(merged,current)


fmt.Println(merged)


}


func main(){
intervals := [][2]int{{1, 3}, {8, 10}, {15, 18} ,{2, 6}}

merge(intervals)

}