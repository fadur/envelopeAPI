package main

import (
	"encoding/json"
	models "envelopeApi/code/src/models"
	utils "envelopeApi/code/src/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*

curl -H 'X-Client-Id: INSERT_CLIENT_ID' \
    -H 'X-Client-Secret: INSERT_CLIENT_SECRET' \
    -H 'Authorization: Bearer ACCESS_TOKEN_HERE' \
    https://api.nordicapigateway.com/v1/accounts
*/

var client, secret string

const (
	url            = "https://api.nordicapigateway.com/v1/accounts"
	unattended_url = "https://api.nordicapigateway.com/v1/authentication/unattended"
	userHash       = "test-user-id"
)

func fetchAccounts(loginToken string) (*models.AccountPayload, error) {

	client	= os.Getenv("CLIENT")
	secret = os.Getenv("SECRET")
	httpClient := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	auth := "Bearer " + loginToken
	fmt.Println(auth)
	req.Header.Add("X-Client-Id", client)
	req.Header.Add("X-Client-Secret", secret)
	req.Header.Add(
		"Authorization",
		auth,
	)
	req.Header.Add("Content-Type", " application/json")
	resp, err := httpClient.Do(req)
	fmt.Println(resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var payload models.AccountPayload
	err = json.Unmarshal([]byte(data), &payload)
	return &payload, err
}

func main() {
	db, err := models.InitDB()
	defer db.Close()
	accessToken := utils.GetToken(db)

	accounts, err := fetchAccounts(accessToken)
	accounts.Save(db)
	if err != nil {
		panic(err)
	}

}
