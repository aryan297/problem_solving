package main

import ("strings" 
        "fmt" )

func lengths(str string){

	word := strings.Fields(str)

	value :=word[len(word)-1]

	fmt.Println(len(value))




}


func main(){
  str := "Hello World"
  lengths(str)
}