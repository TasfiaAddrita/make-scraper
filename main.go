// callback order
// Visit -> OnRequest -> OnResponse -> OnHTML or OnError -> (repeat steps 2-4) -> OnScraped

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

// type job struct {
// 	Company        string `json:"company"`
// 	Title          string `json:"title"`
// 	Location       string `json:"location"`
// 	Department     string `json:"department"`
// 	EmploymentType string `json:"type"`
// 	Description    string `json:"description"`
// }

type Job struct {
	Company string `json:"company"`
	Details struct {
		Title       string `json:"title"`
		Location    string `json:"location"`
		Department  string `json:"department"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"job"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	sweJob := Job{Company: "Brilliant"}

	titleSelector := ".posting-headline > h2"
	employmentTypeSelector := "div.posting-categories"
	descriptionSelector := "div.content-wrapper.posting-page > div > div:nth-child(2)"

	// first callback
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		// fmt.Println("Something went wrong:", err)
	})

	// second callback
	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Visited", r.Request.URL)
	})

	// third-one callback
	c.OnHTML(titleSelector, func(e *colly.HTMLElement) {
		sweJob.Details.Title = e.Text
	})

	// third-two callback
	c.OnHTML(employmentTypeSelector, func(e *colly.HTMLElement) {
		categories := strings.Split(e.Text, "/")
		sweJob.Details.Location = categories[0]
		sweJob.Details.Department = categories[1]
		sweJob.Details.Type = categories[2]
	})

	// third-three callback
	c.OnHTML(descriptionSelector, func(e *colly.HTMLElement) {
		sweJob.Details.Description = e.Text
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://jobs.lever.co/brilliant/359b4cd8-1641-49d0-856e-d457aaa90b01")

	fmt.Printf("%s is hiring a %s %s for their %s location", sweJob.Company, sweJob.Details.Type, sweJob.Details.Title, sweJob.Details.Location)

	// jobJSON, _ := json.Marshal(sweJob)
	// fmt.Println(string(jobJSON))

	file, _ := json.MarshalIndent(sweJob, "", " ")

	_ = ioutil.WriteFile("output.json", file, 0644)

}
