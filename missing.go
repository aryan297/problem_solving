package main


func missing(arr []int){

	left, right:=0 , len(arr)-1


	for left <right{

		mid:=left+(right-left)/2

		if arr[mid]==mid+1{
			left=mid+1
		} else {
			right=mid-1
		}

	}

	return left+1
}


func main(){
 
 arr :=[]int{1,2,4,5}
}