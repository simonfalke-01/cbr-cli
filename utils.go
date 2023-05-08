package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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

func askOpenBrowser(url string) {
	fmt.Print("Open browser? [Y/n] ")
	var answer string
	response, err := fmt.Scanln(&answer)
	if err != nil {
		response = 0
	}

	if response == 0 || answer == "Y" || answer == "y" {
		openBrowser(url)
	} else {
		fmt.Println("The following is the URL to the submission:")
		fmt.Println(url)
	}
}

func checkArgs(num int) {
	if len(os.Args) < num {
		fmt.Println("Usage: ./cbr <problem-id> <path-to-solution>")
		os.Exit(1)
	}
}

func checkIfDarwin() bool {
	return runtime.GOOS == "darwin"
}
