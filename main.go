package main

import (
	"errors"
	"fmt"
	"net/http"
)

type resultRequest struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Fail")

func main() {
	results := make(map[string]string)
	c := make(chan resultRequest)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

// go lang standard library 참조
// chan <-를 함으로써 이 channel은 보내기만 가능해짐 (구체적으로 명시하기 위해)
func hitURL(url string, c chan<- resultRequest) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- resultRequest{url: url, status: "FAILED"}
	} else {
		c <- resultRequest{url: url, status: "OK"}
	}
}
