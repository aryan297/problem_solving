package main

import "fmt"

func rearrange(arr []int){

	positive:=[]int{}
	neg:=[]int{}

	for i:=0;i<len(arr);i++{
		if arr[i]>0{
			positive=append(positive,arr[i])
		}else{
			neg=append(neg,arr[i])
		}
	}

	res:=[]int{}
	i:=0
	j:=0;
    for i<len(positive) && j<len(neg){
          res=append(res,positive[i])
          res=append(res,neg[i])
          i++;
          j++;

    }

       for i < len(positive) {
        res = append(res, positive[i])
        i++
    }

    // Add remaining negatives
    for j < len(neg) {
        res = append(res, neg[j])
        j++
  }
    fmt.Println(res)
	

}


func main(){

	arr:=[]int{1,2,-4,-5,8}
	rearrange(arr)
}