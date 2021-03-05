package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func CreateJobEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var job Job
	json.NewDecoder(request.Body).Decode(&job)
	collection := client.Database(os.Getenv("DB_NAME")).Collection("jobs")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, job)
	err := json.NewEncoder(response).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func GetJobByIDEndpoint(response http.ResponseWriter, request *http.Request) {

}
