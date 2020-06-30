package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	myChannel := make(chan string)

	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.googleddd.com",
		"https://hangouts.googleddd.com",
	}

	for _, site := range sites {

		go getURL(ctx, site, myChannel)

		result := <-myChannel

		if result == "Wrong URL" {
			fmt.Println(result)
			break
		} else {
			fmt.Println(result)
		}

	}

}

func getURL(ctx context.Context, site string, myChannel chan string) {

	res, err := http.Get(site)

	if err != nil {
		myChannel <- "Wrong URL"
	} else {
		myChannel <- res.Status + " " + site
	}

}
