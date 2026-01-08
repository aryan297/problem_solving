package main

import "fmt"


func anagramValue(txt , pat string){

	n:=len(txt)
	k:=len(pat)

	frqMap:=make(map[byte]int)

    ans:=0
	count:=0

	for i:=0;i<len(pat);i++{

		frqMap[pat[i]]++
	}

	count=len(frqMap)

	i,j:=0,0

	for j<n{

		if val,ok:=frqMap[txt[j]];ok{
		    frqMap[txt[j]] = val - 1
		     if frqMap[txt[j]]==0{
		     	count--
		     }

		}

			if j-i+1 < k{
				j++
			} else if j-i+1==k{


				if count==0{
					ans++;
				}

				if val,ok := frqMap[txt[i]];ok{

					frqMap[txt[i]]=val+1;

					if 	frqMap[txt[i]]==1{
						count++
					}


				}

			i++;

		    j++;


			}
     



	}

	fmt.Println(ans)



}


func main(){
txt := "forxxorfxdofr"
pat := "for"

anagramValue(txt,pat)
}