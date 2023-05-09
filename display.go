package main

import "fmt"

func clearLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Printf("\033[1A\x1b[2K")
	}
}

func displayVerdict(verdict Verdict) int {
	fmt.Printf("Verdict:\n")

	lines := 1
	// Display the Subtasks in the side view
	for subtaskNum, subtaskTests := range verdict {
		fmt.Printf("Subtask %d:\n", subtaskNum)
		lines++
		for _, test := range subtaskTests {
			fmt.Printf("- Test Case %s:\n", test.ID)
			fmt.Printf("    Score: %s\n", test.Score)
			fmt.Printf("    Time: %s\n", test.Time)
			fmt.Printf("    Memory: %s\n", test.Memory)
			fmt.Printf("    Verdict: %s\n", test.Verdict)
			lines += 5
		}
		fmt.Printf("\n")
		lines++
	}

	return lines
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
