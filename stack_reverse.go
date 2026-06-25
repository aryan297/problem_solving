package main

import ("fmt"
	    "strconv")

func evaluate( tokens []string){

	stack := []int{}

	for _ , s := range tokens{

		switch s {
		case "+":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack=append(stack,a+b)

		case "*":
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack =stack[:len(stack)-1]

			stack=append(stack,a*b)
			
		default:
			num,_ :=strconv.Atoi(s)
			stack=append(stack,num)

			
		}
	}

	fmt.Println(stack[0] ,"sss")




}


func main() {
	tokens := []string{"2","1","+","3","*"}
	evaluate(tokens)
}