package main

import ("fmt"
	    "strings")

func reverseData(s string){
	s=strings.TrimSpace(s)



	res:=[]string{}
	i:=len(s)-1

	for i>=0{

		for i>=0 && s[i] == ' ' {
			i--
		}
		if i<0{
			break;

		}
		j:=i

		for i>=0 && s[i] != ' '{
			i--
		}

		res=append(res,s[i+1:j+1])
	}

fmt.Println(strings.Join(res, " "))


}


func main(){
	s :="the new town"

	reverseData(s)
}