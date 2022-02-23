package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://pkg.go.dev",
		"https://amazon.com",
	}

	// create a channel to manage child go routines
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Loof for eternity
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
