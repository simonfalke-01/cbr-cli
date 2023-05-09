package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
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

	verdict := Verdict{}
	subtask := 1
	// if testCase is 1, put under new subtask
	for i := range testCases {
		if testCases[i].ID == "1" {
			subtask++
		}

		verdict[subtask] = []TestCase{}
		verdict[subtask] = append(verdict[subtask], testCases[i])
	}

	return verdict
}
