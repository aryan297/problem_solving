package main

import "fmt"


func subset(nums []int , target int) bool{

	n:=len(nums)

	dp:=make([][]bool , n+1)

	for i := range dp{
		dp[i]=make([]bool , target+1)
		dp[i][0]=true
	}

	for i:=1;i<=n;i++{
		for j:=1;j<=target;j++{

			if nums[i-1]<=j{
				dp[i][j]=dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else{
				dp[i][j]=dp[i-1][j]
			}

		}
	}

   return dp[n][target]

}


func main(){
	nums :=[]int{3,4,5}

	k :=subset(nums,10)

	fmt.Println(k)
}