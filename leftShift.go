package main


import "fmt"

func leftArray(s , goal string) bool{

	rooted :=[]byte(s)

	fmt.Println(rooted)

	n:=len(rooted)

	for i:=0;i<n;i++{
		//fmt.Println(s[i])

		if string(rooted) == goal{
			return true
		}

		first:=rooted[0]

		fmt.Println(string(first) ,"k")

		for j:=0;j<n-1;j++{
			rooted[j] =rooted[j+1]

		}

		rooted[n-1]=first

		fmt.Println(string(first),"h")


	}

	fmt.Println(goal)

	return false







}

func main(){
   s := "abcde"
goal := "cdeab"

 fmt.Println(leftArray(s, goal))

}