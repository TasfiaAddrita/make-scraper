// callback order
// Visit -> OnRequest -> OnResponse -> OnHTML or OnError -> (repeat steps 2-4) -> OnScraped

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Job struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	URL            string             `json:"url" bson:"url"`
	Company        string             `json:"company" bson:"company"`
	Title          string             `json:"title" bson:"title"`
	Location       string             `json:"location" bson:"location"`
	Department     string             `json:"department" bson:"department"`
	EmploymentType string             `json:"employmentType" bson:"employmentType"`
	Description    string             `json:"description" bson:"description"`
}

func getJobDetails(c *colly.Collector, jobLink string) Job {
	// c.Visit("https://jobs.lever.co/brilliant/359b4cd8-1641-49d0-856e-d457aaa90b01")
	c.Visit(jobLink)

	job := Job{Company: "Brilliant"}

	titleSelector := ".posting-headline > h2"
	employmentTypeSelector := "div.posting-categories"
	descriptionSelector := "div.content-wrapper.posting-page > div > div:nth-child(2)"

	c.OnRequest(func(r *colly.Request) {})

	c.OnResponse(func(r *colly.Response) {})

	c.OnHTML(descriptionSelector, func(e *colly.HTMLElement) {
		job.Description = e.Text
	})

	c.OnHTML(titleSelector, func(e *colly.HTMLElement) {
		job.Title = e.Text
	})

	c.OnHTML(employmentTypeSelector, func(e *colly.HTMLElement) {
		categories := strings.Split(e.Text, "/")
		job.Location = categories[0]
		job.Department = categories[1]
		job.EmploymentType = categories[2]
	})

	c.OnError(func(_ *colly.Response, err error) {})

	c.OnScraped(func(r *colly.Response) {})

	fmt.Printf("%s is hiring a %s %s for their %s location", job.Company, job.EmploymentType, job.Title, job.Location)

	return job
}

func main() {

	// ---------FLAG--------------
	// var jobLinkFromCLI string

	// flag.StringVar(&jobLinkFromCLI, "url", "", "Add a job from URL.")

	// if jobLinkFromCLI != "" {
	// getJobLinkFromCLI() > getJobDetails > addJobToDB (POST request)
	// }

	// Instantiate default collector
	// c := colly.NewCollector()

	// getJobDetails(c, jobLinkFromCLI)

	fmt.Println("Starting application...")
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	fmt.Println("Successfully connected to DB...")

	router := mux.NewRouter()
	router.HandleFunc("/api/jobs", CreateJobEndpoint).Methods("POST")
	router.HandleFunc("/api/jobs", GetAllJobsEndpoint).Methods("GET")
	http.ListenAndServe(":5000", router)

}
