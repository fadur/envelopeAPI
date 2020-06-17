package main

import (
	"bytes"
	"context"
	"encoding/json"
	models "envelopeApi/code/src/models"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
	"os"
)

var client, secret string

const (
	url    = "https://api.nordicapigateway.com/v1/authentication/tokens"

)

func authCode(code string) (map[string]interface{}, error) {

	client	= os.Getenv("CLIENT")
	secret  = os.Getenv("SECRET")
	httpClient := &http.Client{}
	body := map[string]string{"code": code}
	payload, err := json.Marshal(&body)
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-Client-Id", client)
	req.Header.Add("X-Client-Secret", secret)
	req.Header.Add("Content-Type", " application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	var objmap map[string]interface{}
	json.Unmarshal([]byte(data), &objmap)
	if err != nil {
		panic(err)
	}
	return objmap, nil
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := models.InitDB()
	var data models.Payload
	defer db.Close()
	code := req.QueryStringParameters["code"]
	payload, err := authCode(code)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := json.Marshal(&payload)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(resp, &data)
	fmt.Printf("%v\n", data)
	data.Save(db)
	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}
