package main


import ("fmt"
        "sort")

func merge(arr1 ,arr2 []int , m ,n int){
 i:=m-1;
 j:=n-1
 k:=m+n-1

 for  j>=0{


     if i>=0 && arr1[i]>arr2[j]{

     	arr1[k]=arr1[i]
     	i--
     } else{
     	arr1[k]=arr2[j]
     	j--
     }
     k--
 }

fmt.Println(arr1)

}


func main(){

	arr1 :=[]int{-5, -2, 4, 5, 0, 0, 0 ,0 ,0 ,0,0}
	arr2 := []int{5, -3, -2, 1, 4, 5, 8}


	validArray1 := 4
	validArray2:=7
		// Sort both arrays before merging
	sort.Ints(arr1[:validArray1])
	sort.Ints(arr2)

	merge(arr1,arr2 , validArray1, validArray2)
}