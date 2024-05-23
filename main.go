package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

var errRequestFailed = errors.New("요청에 실패했습니다")

func main() {
	results := make(map[string]string) // initailize empty map
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		go checkUrl(url, c)
	}
	
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

//  chan<- channel send only
func checkUrl(url string, c chan<- requestResult)  {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url:url, status: "FAILED"}
	}
	c <- requestResult{url:url, status: status}
	
}