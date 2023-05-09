package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Attempt struct {
	ID         string `json:"id"`
	Submission string `json:"submission"`
	Username   string `json:"username"`
	Problem    string `json:"problem"`
	Score      string `json:"score"`
	Language   string `json:"language"`
	MaxTime    string `json:"max_time"`
	MaxMemory  string `json:"max_memory"`
}

type TestCase struct {
	ID      string `json:"id"`
	Score   string `json:"score"`
	Verdict string `json:"verdict"`
	Time    string `json:"time"`
	Memory  string `json:"memory"`
}

type Verdict map[int][]TestCase

func processString(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func askToOpenBrowser(url string) {
	fmt.Print("Would you like to open browser to view the submission? [y/N] ")
	var answer string
	response, err := fmt.Scanln(&answer)
	if err != nil {
		response = 0
	}

	if response == 0 || answer == "N" || answer == "n" {
		fmt.Println("The following is the URL to the submission:")
		fmt.Println(url)
	} else {
		openBrowser(url)
	}
}

func checkArgs(num int) {
	if len(os.Args) < num {
		fmt.Println("Usage: ./cbr <problem-id> <path-to-solution>")
		os.Exit(1)
	}
}

func systemIsDarwin() bool {
	switch runtime.GOOS {
	case "darwin":
		return true
	default:
		return false
	}
}

func sleep(seconds int) {
	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	for {
		select {
		case <-ticker.C:
			return
		}
	}
}
