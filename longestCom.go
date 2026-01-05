package main

import "fmt"

func longestCommonPrefix(s []string)string{

	if len(s)==0{
		return ""
	}

	prefix:=s[0]

	for _, word :=range s[1:]{

		for len(word)<len(prefix) || word[:len(prefix)] != prefix{

			prefix=prefix[:len(prefix)-1]

			if prefix==""{
				return ""
			}
		}

	}

	return prefix


}




func main() {

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"


}