package main

import "fmt"

func main() {

	var username string = "maverick"
	var email string
	email = "maverick@example.com"

	fmt.Printf("username : %s\nemail : %s\n", username, email)

	city := "Jakarta"
	fmt.Printf("city : %s\n",city)

	var golang, java, php string
	golang, java, php = "golang","java","php"
	fmt.Printf("skills : %s, %s, dan %s\n", golang, java, php)

	orange, apple, mango := "orange", "apple", "mango"
	fmt.Printf("most popular fruits : %s, %s, and %s\n", orange, apple, mango)

	var music, _ string = "grunge", "rock"
	fmt.Printf("fav music : %s", music)
}
