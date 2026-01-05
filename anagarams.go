package main

import ("fmt" 
	    "strings" 
	    "sort")

func anagarm(s []string){

	hashMap:=make(map[string][]string)

	for _ , a := range s{

		ch:=strings.Split(a,"")

		sort.Strings(ch)

		key:= strings.Join(ch,"")



		hashMap[key]=append(hashMap[key],a)

	}


	res:=[][]string{}

	for _ , data := range hashMap{
		res=append(res,data)

	}


	fmt.Println(res)

}


func main(){
	 input := []string{"abb", "bba", "abc", "cab", "bca"}

	 anagarm(input)
}