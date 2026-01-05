package main

import "fmt"


func isomorphic(s string , t string) bool{
if len(s)!=len(t){
	return false ;
}
  mapS:=make(map[byte]byte)
  mapT:=make(map[byte]byte)

  for i:=0;i<len(s);i++{

  	c1,c2:=s[i],t[i]

  	if val, ok := mapS[c1];ok{

  		if val!=c2{
  			return false
  		} 
  		}else{
  			mapS[c1]=c2


  	}

   if val , ok :=mapT[c2];ok{

   	  if val!=c1{
   	  	return false
   	  } 
   	  }else{
         mapT[c2]=c1
   }


  }

return true

}


func main() {
	s:="paper"
	t:="title"

	fmt.Println(isomorphic(s,t))
	fmt.Println(isomorphic("foo", "bar")) 
}