package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func AddJobToDB(job Job) (*mongo.InsertOneResult, error) {
	collection := client.Database(os.Getenv("DB_NAME")).Collection("jobs")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, job)
	if err != nil {
		fmt.Println("Could not add job to database.")
	} else {
		fmt.Println("Added job to database.")
	}
	return result, err
}

// CreateJobEndpoint POST api/jobs
func CreateJobEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var job Job
	json.NewDecoder(request.Body).Decode(&job)
	result, _ := AddJobToDB(job)
	json.NewEncoder(response).Encode(result)
}

// GetAllJobsEndpoint GET api/jobs
func GetAllJobsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var jobs []Job
	collection := client.Database(os.Getenv("DB_NAME")).Collection("jobs")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var job Job
		cursor.Decode(&job)
		jobs = append(jobs, job)
		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
	}

	json.NewEncoder(response).Encode(jobs)
	fmt.Println("GET ALL endpoint being called")
}
