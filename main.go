// callback order
// Visit -> OnRequest -> OnResponse -> OnHTML or OnError -> (repeat steps 2-4) -> OnScraped

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Job struct {
	URL            string `json:"url"`
	Company        string `json:"company"`
	Title          string `json:"title"`
	Location       string `json:"location"`
	Department     string `json:"department"`
	EmploymentType string `json:"type"`
	Description    string `json:"description"`
}

func scrapeWebsite(c *colly.Collector, jobLink string) {
	c.Visit("https://jobs.lever.co/brilliant/359b4cd8-1641-49d0-856e-d457aaa90b01")

	sweJob := Job{Company: "Brilliant"}

	titleSelector := ".posting-headline > h2"
	employmentTypeSelector := "div.posting-categories"
	descriptionSelector := "div.content-wrapper.posting-page > div > div:nth-child(2)"

	c.OnRequest(func(r *colly.Request) {})

	c.OnResponse(func(r *colly.Response) {})

	c.OnHTML(descriptionSelector, func(e *colly.HTMLElement) {
		sweJob.Description = e.Text
	})

	c.OnHTML(titleSelector, func(e *colly.HTMLElement) {
		sweJob.Title = e.Text
	})

	c.OnHTML(employmentTypeSelector, func(e *colly.HTMLElement) {
		categories := strings.Split(e.Text, "/")
		sweJob.Location = categories[0]
		sweJob.Department = categories[1]
		sweJob.EmploymentType = categories[2]
	})

	c.OnError(func(_ *colly.Response, err error) {})

	c.OnScraped(func(r *colly.Response) {})

	fmt.Printf("%s is hiring a %s %s for their %s location", sweJob.Company, sweJob.EmploymentType, sweJob.Title, sweJob.Location)

	// jobJSON, _ := json.Marshal(sweJob)
	// fmt.Println(string(jobJSON))

	// file, _ := json.MarshalIndent(sweJob, "", " ")
	// _ = ioutil.WriteFile("output.json", file, 0644)
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	// ---------FLAG--------------
	var jobLinkFromCLI string

	flag.StringVar(&jobLinkFromCLI, "url", "", "Add a job from URL.")

	if jobLinkFromCLI != "" {
		// getJobLinkFromCLI() > scrapeWebsite > addJobToDB (POST request)
	}

	// Instantiate default collector
	c := colly.NewCollector()

	scrapeWebsite(c, jobLinkFromCLI)

}
