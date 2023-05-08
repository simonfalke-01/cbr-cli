package main

import (
	"bytes"
	"fmt"
	"github.com/simonfalke-01/cbr-cli/kooky"
	_ "github.com/simonfalke-01/cbr-cli/kooky/browser/all"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`codebreaker.xyz`), kooky.Name(`google-login-session`))

	cookie := cookies[0].Value

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <problem name> <solution path>")
		os.Exit(1)
	}

	problemName := os.Args[1]
	urlStr := fmt.Sprintf("https://codebreaker.xyz/problem/%s", problemName)

	solutionPath := os.Args[2]
	// read solutionPath
	code, err := os.ReadFile(solutionPath)
	if err != nil {
		panic(err)
	}

	formData := url.Values{
		"language": {"C++ 17"},
		"code":     {string(code)},
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "testCookie; google-login-session="+cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/111.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Origin", "https://codebreaker.xyz")
	req.Header.Set("Referer", urlStr)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Te", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	fmt.Println("Submitted.")
}
