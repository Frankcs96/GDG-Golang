/*
Console application.

This application simulates a log in with 3 hard coded accounts

Account---Password
user1	  1234
user2     1234
user3     1234

If you log in, your account will be cached so  if you try to enter a second time
you won't have to enter your password again.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fcsuarez96/http/user"
)

const (
	//URL from server
	URL = "http://localhost:8080"
)

func main() {

	client := newClient()

	running := true

	fmt.Println("Welcome to GDG Marbella!\n===========================")
	for running {
		fmt.Println("Enter a number!")
		fmt.Println("1.Log in")
		fmt.Println("2.Exit")
		var option string
		fmt.Scanln(&option)

		if option == "1" {

			fmt.Println("Enter your username: ")
			var username string
			fmt.Scanln(&username)
			//Check if user is in cache so we can skip password

			userWithoutPassword := user.NewUser(username, "")

			buffer := new(bytes.Buffer)

			json.NewEncoder(buffer).Encode(userWithoutPassword)

			resp, err := client.Post(URL+"/cache", "application/json; charset=utf-8", buffer)

			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Println("Welcome Back " + username)

			} else {

				//If not in cache continue with password

				fmt.Println("Enter your password: ")
				var password string
				fmt.Scanln(&password)

				user := user.NewUser(username, password)

				buf := new(bytes.Buffer)
				json.NewEncoder(buf).Encode(user)

				resp, err = client.Post(URL+"/login", "application/json; charset=utf-8", buf)

				if err != nil {
					fmt.Println(err)
				}
				defer resp.Body.Close()

				if resp.StatusCode == 200 {
					fmt.Println("Welcome " + username)
				} else {
					fmt.Println("Wrong user please try again!")
				}
			}

		}

		if option == "2" {
			running = false
		}
	}

}
