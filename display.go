package main

import (
	"fmt"
	"os"
	"os/exec"
)

func lessDisplay(s string) {
	cmdStr := fmt.Sprintf("echo \"%s\" | less", s)
	cmd := exec.Command("bash", "-c", cmdStr, "-R")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func displayVerdict(verdict Verdict) {
	// display all non-ac testcases for each subtask
	ac := true
	var lessDisplayStr string
	for i, subtask := range verdict {
		printed := false
		for _, testCase := range subtask {
			if testCase.Verdict != "AC" {
				ac = false
				if !printed {
					lessDisplayStr += fmt.Sprintf("Subtask %d:\n", i)
					printed = true
				}
				lessDisplayStr += fmt.Sprintf("- Test Case %s:\n", testCase.ID)
				lessDisplayStr += fmt.Sprintf("    Score: %s\n", testCase.Score)
				lessDisplayStr += fmt.Sprintf("    Time: %s\n", testCase.Time)
				lessDisplayStr += fmt.Sprintf("    Memory: %s\n", testCase.Memory)
				lessDisplayStr += fmt.Sprintf("    Verdict: %s\n", testCase.Verdict)
			}
		}
	}

	if !ac {
		lessDisplay(lessDisplayStr)
	} else {
		fmt.Println("Hurray! AC!")
	}
}

func hasIncomplete(verdict Verdict) bool {
	for _, subtaskTests := range verdict {
		for _, test := range subtaskTests {
			if test.Verdict == ":(" {
				return true
			}
		}
	}
	return false
}
