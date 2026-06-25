package main

import ("fmt")

func  stacks(str string){

	stack :=[]byte{}

	for _,  k := range str{
		ch:=byte(k)
		if len(stack)>0 && stack[len(stack)-1]==ch{
			stack=stack[:len(stack)-1]
		} else {
			stack=append(stack,ch)
		}
	}

	fmt.Println(string(stack))
}

func main(){
	str :="abbaca"
	stacks(str)
}