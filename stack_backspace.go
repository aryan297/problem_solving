package main

import "fmt"


func backSpace(s string) string{

	stack :=[]byte{}

	for _ , k := range s {

		ch :=byte(k)

		if ch !='#'{
			stack=append(stack,ch)
		} else if len(stack)>0{
			stack=stack[:len(stack)-1]
		}
	}

	return string(stack)

}

func buildStack(s , t string) bool{
	return backSpace(s)==backSpace(t)
}




func main(){
s := "ab##"
t := "c#d#"

b := buildStack(s,t)

fmt.Println(b)
}