package main

import (
	"fmt"
	"time"
)

/*
import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Fail")

func main() {
	var results = make(map[string]string)
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
	results["hello"] = "Hello"
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	fmt.Println(results)

}

// go lang standard library 참조
func hitURL(url string) error {
	fmt.Println("체크", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
*/

func main() {
	c := make(chan string) //channel과 어떤 타입을 보낼건지
	people := [6]string{"신", "홍", "섭", "개", "발", "자"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Print(i)
		fmt.Println(<-c) //<-c는 blocking operation
	}
}

//Channel을 이용하여 true 메세지를 보낸다, c로 true메세지를 보낸다
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person
}
