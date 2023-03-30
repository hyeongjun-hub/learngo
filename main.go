package main // compile 하고싶으면 무조건 main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

type requestStruct struct {
	url    string
	status int
}

func main() {
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

	results := make(map[string]int)
	c := make(chan requestStruct)

	for _, url := range urls {
		// go routine 사용
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		tempStruct := <-c
		results[tempStruct.url] = tempStruct.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

// hitURL: chan<- 로 channel 을 이 func 에서 send-only 로 설정
func hitURL(url string, c chan<- requestStruct) {
	//fmt.Println("Checking", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		// error handle
		fmt.Println("Error!")
	}
	c <- requestStruct{url, resp.StatusCode}
}

// python 처럼 하나씩 말고 go를 이용해 동시에 한방에 하도록
// goroutine 사용 !
