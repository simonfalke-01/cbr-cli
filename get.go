package main

import (
	"io"
	"net/http"
)

func getPage(url string) io.ReadCloser {
	get, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	return get.Body
}
