package main


import "fmt"

func groupAnagram(s []string) {
	hashMap := make(map[[26]int][]string)

	for _ , str := range s{

		var count [26]int

		for i:=0;i<len(str); i++{

			count[str[i]-'a']++

		}
		hashMap[count]=append(hashMap[count],str)
	
	}

res := make([][]string,0, len(hashMap))


		for _, group := range hashMap{

			res=append(res,group)
		}

	fmt.Println(res)


}

func main(){

	str := []string{"eat","tea","tan","ate","nat","bat"}

	groupAnagram(str)


}