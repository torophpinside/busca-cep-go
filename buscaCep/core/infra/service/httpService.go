package service

import (
	"net/http"
)

func Call(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		resp.Body.Close()
		return nil
	}

	return resp
}
