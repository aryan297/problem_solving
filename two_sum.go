 package main

 import "fmt"

 func twoSum(nums []int , target int) []int{

 	seen := make(map[int]int)

 	for i , n := range nums {
 		complement :=target -n

 		if j , ok := seen[complement];ok{

 			return []int{j ,i}

 		}

 		seen[n]=i
 	}

 	return nil
 }


 func main() {

 	nums := []int{2,7,11,15}
 	target :=18

 	result := twoSum(nums,target)

 	fmt.Println(result)
 }