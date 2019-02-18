package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
)

const (
	url    = "https://api.nordicapigateway.com/v1/authentication/tokens"
	secret = "043a97a8f3d27a749b6cde568737dbe36b0ef1c273683cd4ffd7f09c25989f38"
	client = "fadur-9ad89068-c4ec-4c98-afdf-631e177415f6"
)

func authCode(code string) (map[string]interface{}, error) {
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
	code := req.QueryStringParameters["code"]
	payload, err := authCode(code)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := json.Marshal(&payload)
	if err != nil {
		fmt.Println(err)
	}
	var data Payload
	json.Unmarshal(resp, &data)
	fmt.Println(data.Login.Label)
	fmt.Println(data.Login.Expires)
	fmt.Println(data.Session.Expires)

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}
