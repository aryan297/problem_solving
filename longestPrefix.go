package main

import "fmt"


func prefixData(s []string){
    if len(s)==0{
    	fmt.Println("No data")
    }

    prefix :=s[0]


    for _, str :=range s[1:]{

    	    i:=0

    	for  i < len(prefix) && i < len(str) && prefix[i]==str[i] {
    		i++
    	}

    	prefix=prefix[:i]

    if prefix == ""{
    	break;
      }
    
    }




    fmt.Println(prefix)


}


func main(){

	str:=[]string{"flower","fly","flow","flight"}

	prefixData(str)

}