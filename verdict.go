package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"sort"
	"strconv"
)

func parseVerdict(reader io.ReadCloser) Verdict {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}

	var testCases []TestCase

	doc.Find(".table tbody tr").Each(func(i int, s *goquery.Selection) {
		testCase := TestCase{}

		s.Find("td").Each(func(j int, ss *goquery.Selection) {
			switch j {
			case 0:
				testCase.ID = processString(ss.Text())
			case 1:
				testCase.Score = processString(ss.Text())
			case 2:
				testCase.Verdict = processString(ss.Text())
			case 3:
				testCase.Time = processString(ss.Text())
			case 4:
				testCase.Memory = processString(ss.Text())
			}
		})

		testCases = append(testCases, testCase)
	})

	testCases = testCases[6:]

	verdict := Verdict{}
	subtask := 1
	// if testCase is 1, put under new subtask
	for i := range testCases {
		if testCases[i].ID == "1" {
			subtask++
			verdict[subtask] = []TestCase{}
		}

		verdict[subtask] = append(verdict[subtask], testCases[i])
	}

	sortVerdictByKeys(verdict)

	return verdict
}

func getVerdict(submissionID int) Verdict {
	resp := getPage(fmt.Sprintf("https://codebreaker.xyz/submission/%s", strconv.Itoa(submissionID)))
	verdict := parseVerdict(resp)
	incomplete := hasIncomplete(verdict)

	if incomplete {
		sleep(200)
		return getVerdict(submissionID)
	} else {
		return verdict
	}
}

func sortVerdictByKeys(verdict Verdict) Verdict {
	// Extract the keys from the map
	keys := make([]int, 0, len(verdict))
	for key := range verdict {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Ints(keys)

	// Create a new sorted Verdict map
	sortedVerdict := make(Verdict)
	for _, key := range keys {
		sortedVerdict[key] = verdict[key]
	}

	return sortedVerdict
}
