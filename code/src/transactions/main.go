package main

import (
	// 	"bytes"
	"encoding/json"
	models "envelopeApi/code/src/models"
	"envelopeApi/code/src/utils"
	"fmt"
	// "github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
)

const (

	client	= os.Getenv('CLIENT')
	secret = os.Getenv('SECRET')
	base_url       = "https://api.nordicapigateway.com/v1/accounts"
	unattended_url = "https://api.nordicapigateway.com/v1/authentication/unattended"
	userHash       = "test-user-id"
)

/*

curl https://api.nordicapigateway.com/v1/accounts/{accountId}/transactions?FromDate={YYYY-MM-DD} \
  -H 'X-Client-Id: CLIENT_ID' \
  -H 'X-Client-Secret: CLIENT_SECRET' \
  -H 'Authorization: Bearer ACCESS_TOKEN'
*/

func fetchTransactions(account string, accessToken string) (*models.TransactionResponse, error) {
	uri := fmt.Sprintf("%s/%s/transactions?fromDate=2019-01-01", base_url, account)
	fmt.Println(uri)
	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		uri,
		nil,
	)
	req.Header.Add("X-Client-Id", client)
	req.Header.Add("X-Client-Secret", secret)
	req.Header.Add("Content-Type", " application/json")
	req.Header.Add(
		"Authorization",
		fmt.Sprintf("Bearer %s", accessToken),
	)
	resp, err := httpClient.Do(req)
	fmt.Println(resp.StatusCode)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var payload models.TransactionResponse
	err = json.Unmarshal([]byte(data), &payload)
	if err != nil {
		fmt.Println(err)
	}
	return &payload, err
}

func main() {
	db, err := models.InitDB()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	accessToken := utils.GetToken(db)
	var accountIds []string
	db.Table("accounts").Pluck("id", &accountIds)
	for _, account := range accountIds {
		transactionResponse, err := fetchTransactions(account, accessToken)
		if err != nil {
			fmt.Println(err)
		}
		transactionResponse.Save(db)
		fmt.Println("save called")

	}
}
