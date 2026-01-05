package main 

import "fmt"


func star1(){
	for i:=0;i<5;i++{
		for j:=0; j<5;j++{

			fmt.Print("*")
		}

		fmt.Println("")


	}
}


func main(){
	star1()

}