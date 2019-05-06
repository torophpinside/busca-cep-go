package service

import (
	"net/http"
)

var channelResponse chan (*http.Response) = make(chan *http.Response)

func Call(url string) *http.Response {
	go fromUrl(url)
	return <-channelResponse
}

func fromUrl(url string) {
	resp, _ := http.Get(url)
	channelResponse <- resp
}
