package main


import ("fmt"
        "math")


func largestArray(arr []int) (int, int, int){
	maxFirst:=math.MinInt
	maxSecond:=math.MinInt
	maxThird:=math.MinInt

	for i:=0;i<len(arr);i++{

		if arr[i]>maxFirst{
			maxThird=maxSecond
			maxSecond=maxFirst
			maxFirst=arr[i]
		} else if arr[i]>maxSecond{

			maxThird=maxSecond
			maxSecond=arr[i]
		} else if arr[i]>maxThird{
			maxThird=arr[i]
		}



	}

	return maxFirst, maxSecond,maxThird
}


func main(){

	arr :=[]int{2,5,1,3,0}
	fmt.Println(largestArray(arr))

}

