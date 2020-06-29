package main

import (
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, site := range sites {
		wg.Add(1)

		go func(site string) {

			res, err := http.Get(site)

			if err != nil {

			}

			io.WriteString(os.Stdout, res.Status+"\n")

			wg.Done()
		}(site)
	}

}

// time sleep
