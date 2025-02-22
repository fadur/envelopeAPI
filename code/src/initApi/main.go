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
	url         = "https://api.nordicapigateway.com/v1/authentication/initialize"
	callbackUrl = "http://localhost:3000/callbackUrl"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	client	= os.Getenv("CLIENT")
	secret = os.Getenv("SECRET")
	db, err := models.InitDB()
	defer db.Close()
	httpClient := &http.Client{}

	body := map[string]string{"userHash": "test-user-id", "redirectUrl": callbackUrl}
	payload, err := json.Marshal(&body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", string(payload))
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
	var objmap map[string]string
	json.Unmarshal([]byte(data), &objmap)
	if err != nil {
		panic(err)
	}
	authUrl := objmap["authUrl"]
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"location": authUrl,
		},
		StatusCode: 302,
	}, nil

}

func main() {
	lambda.Start(handler)
}
