package main

import ("encoding/json"
        "fmt"
        "time"
    )

type User struct{

	Name string `json:"name"`
	Age int     `json:"age"`
	Date time.Time `json:"date"`	
}

func main() {

	jsons :=`{
	"name" :"Raj",
	"age" :21,
    "date": "2026-01-23T10:30:00Z"}
    `


    userData := User{
    	Name: "Raju",
    	Age: 21,
    	Date:time.Now().UTC(),
    }

    var user User

      err :=json.Unmarshal([]byte(jsons),&user)

      data , _ := json.MarshalIndent(userData,""," ")

    if err!=nil{
    	fmt.Println("rr",err)
    	return
    }

    fmt.Println("User:", user)
	fmt.Println("Created At:", user.Date)
	fmt.Println(string(data))

	fmt.Println("Year:", user.Date.Year())



}