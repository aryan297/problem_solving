package main

import "fmt"

type Stack struct{

Items []int

}

func (s *Stack) Push(val int) {
	s.Items=append(s.Items,val)
}

func (s *Stack) Pop() int{
	if len(s.Items)==0{
		return -1
	}
	top := s.Items[len(s.Items)-1]
	s.Items=s.Items[:len(s.Items)-1]
	return top
}



func(s *Stack) Println(){
	for _,k := range s.Items{
		fmt.Println(k)
	}
}

func (s *Stack) IsEmpty() bool{
	if len(s.Items)==0{
		return true
	} 
	return false
}

func ( s *Stack) Peek() int{

	if len(s.Items)==0{
		return -1
	}

    item := s.Items[len(s.Items)-1]
	return item
}


func isValid(s string){
	
}



func main(){
	s := &Stack{}
	s.Push(1)
	s.Push(3)
	s.Push(35)
	s.Push(4)
	k :=s.Pop()
	j:= s.Peek()

	fmt.Println(k ,"pop")
	fmt.Println(j,"eee")

	s.Println()






}