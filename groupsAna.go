package main 

import ("fmt"
        "sort")

func anagram(strs []string){

	groups :=make(map[string][]string)

	for _,v := range strs{

		b :=[]byte(v)

		sort.Slice(b ,func(i, j int ) bool { return b[i]< b[j]})

		key:=string(b)

		groups[key]=append(groups[key],v )

	}

	res := make([][]string,0,len(groups))

	for _, v := range groups{

		res = append(res,v)
	}

	fmt.Println(res)


}


func main() {

	strs :=  []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagram(strs)
}