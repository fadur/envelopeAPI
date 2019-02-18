package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"models"
)

type MyEvent struct {
	Name string `json:"name"`
}

type PeopleResponse struct {
	People []Person `json:"data"`
}

func handler(ctx context.Context, event MyEvent) (PeopleResponse, error) {
	db := models.InitDB()
	defer db.Close()
	people := GetPeople()
	return PeopleResponse{people}, nil
}

func main() {

	lambda.Start(handler)
}
