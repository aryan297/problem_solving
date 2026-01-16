package main

import (
	"fmt"
	"strings"
)

type UserActivity struct {
	login  int
	upload int
	logout int
}

func userInput() {
	userStr := []string{
		"2023-06-18T10:00:00Z user1 login",
		"2023-06-18T10:05:00Z user2 login",
		"2023-06-18T10:10:00Z user1 upload",
		"2023-06-18T10:15:00Z user2 upload",
		"2023-06-18T10:20:00Z user1 logout",
	}

	userActivity := make(map[string]*UserActivity)

	for i := 0; i < len(userStr); i++ {
		parts := strings.Split(userStr[i], " ")
		if len(parts) != 3 {
			continue
		}

		user := parts[1]
		action := parts[2]

		if _, exists := userActivity[user]; !exists {
			userActivity[user] = &UserActivity{}
		}

		switch action {
		case "login":
			userActivity[user].login++
		case "upload":
			userActivity[user].upload++
		case "logout":
			userActivity[user].logout++
		}
	}

	// Output
	fmt.Println("unique_users:", len(userActivity))
	fmt.Println("user_activity:")
	for user, activity := range userActivity {
		fmt.Println(user, "=>", *activity)
	}
}

func main() {
	userInput()
}
