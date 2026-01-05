package main

import "fmt"


func longest( str []string ){

	for i:=0;i<len(str);i++{

		ch:=str[0][i]

		for j:=1;j<len(str);j++{

			if i> len(str) || str[i][j]!=ch{
				fmt.Println(str[0][:i])
			}
		}
	}


}


func main(){

	str :=[]string{"flower","flow","flight"}
	longest(str)

}