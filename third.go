package main

import "fmt"

func findThird(arr []int) {
	if len(arr) < 3 {
		fmt.Println("Array must have at least 3 elements")
		return
	}

	const INF = int(^uint(0) >> 1) // max int value
	first, second, third := INF, INF, INF

	for _, num := range arr {
		if num < first {
			third = second
			second = first
			first = num
		} else if num < second && num != first {
			third = second
			second = num
		} else if num < third && num != second && num != first {
			third = num
		}
	}

	if third == INF {
		fmt.Println("No third smallest element (array may not have enough unique elements)")
	} else {
		fmt.Println("Third smallest:", third)
	}
}

func main() {
	arr := []int{7, 9, 9, 1, 3, 5, 2}
	findThird(arr) // Expected → 3
}
