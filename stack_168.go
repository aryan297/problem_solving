package main

import ("fmt"
        "strconv")

func checkValid(s []string){

	stack :=[]int{}

	for _, ch := range s{

		switch ch {
		case "C":
			 stack=stack[:len(stack)-1]

		case "D":
			last:=stack[len(stack)-1]
			stack=append(stack,last*2)

		case "+":
			sum :=stack[len(stack)-1] + stack[len(stack)-2]
			stack=append(stack,sum)

		
		default:
			num ,_ :=strconv.Atoi(ch)

			stack=append(stack,num)
			

		}
	}

	ans:=0
	for _, v := range stack{

		ans+=v
	}

	fmt.Println(ans)


}



func main(){

 	 ops:= []string{"5", "2", "C", "D", "+"}
 	checkValid(ops)

}