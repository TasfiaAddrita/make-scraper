package main

import "testing"

func TestGetJobDetailsFromURL(t *testing.T) {
	jobLink := "https://jobs.lever.co/lever/45b30441-05e6-4eae-b004-f2f9895f6942"
	job := GetJobDetailsFromURL(jobLink)
	var tests = []struct {
		input    string
		expected string
	}{
		{job.Company, "Lever "},
		{job.URL, jobLink},
		{job.Title, "Associate Customer Success Manager"},
	}

	for _, test := range tests {
		if output := test.input; output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func BenchmarkTestGetJobDetailsFromURL(b *testing.B) {
	jobLink := "https://jobs.lever.co/lever/45b30441-05e6-4eae-b004-f2f9895f6942"
	for i := 0; i < b.N; i++ {
		GetJobDetailsFromURL(jobLink)
	}
}
