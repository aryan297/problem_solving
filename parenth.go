package main


import "fmt"


func parenth(s string){

	var res []byte

	balance:=0

	for i:=0;i<len(s);i++{

		if s[i] == '('{
			fmt.Println(balance)

			if balance>0{
				res=append(res,'(')
			}
			balance++

		} else {
					balance--

			if balance >0{
					res=append(res,')')
			}

	
		}


	}

	fmt.Println(string(res))

}

func main(){

	s:="(()())(())"
	parenth(s)

}