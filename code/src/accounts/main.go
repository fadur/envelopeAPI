package main

import (
	"bytes"
	"encoding/json"
	models "envelopeApi/code/src/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*

curl -H 'X-Client-Id: INSERT_CLIENT_ID' \
    -H 'X-Client-Secret: INSERT_CLIENT_SECRET' \
    -H 'Authorization: Bearer ACCESS_TOKEN_HERE' \
    https://api.nordicapigateway.com/v1/accounts
*/

const (
	secret         = "043a97a8f3d27a749b6cde568737dbe36b0ef1c273683cd4ffd7f09c25989f38"
	client         = "fadur-9ad89068-c4ec-4c98-afdf-631e177415f6"
	url            = "https://api.nordicapigateway.com/v1/accounts"
	unattended_url = "https://api.nordicapigateway.com/v1/authentication/unattended"
	userHash       = "test-user-id"
)

func getSession(token string) (models.Session, error) {

	body := map[string]string{"userHash": userHash, "loginToken": token}
	payload, err := json.Marshal(&body)

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		unattended_url,
		bytes.NewBuffer(payload),
	)
	req.Header.Add("X-Client-Id", client)
	req.Header.Add("X-Client-Secret", secret)
	req.Header.Add("Content-Type", " application/json")

	resp, err := httpClient.Do(req)
	fmt.Println(resp.StatusCode)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	var objmap map[string]interface{}
	err = json.Unmarshal([]byte(data), &objmap)

	if err != nil {
		fmt.Println(err)
	}
	session := objmap["session"].(map[string]interface{})
	accessToken := session["accessToken"].(string)
	fmt.Println(accessToken)
	expires := session["expires"].(string)
	t, err := time.Parse(time.RFC3339, expires)
	fmt.Println(t)
	return models.Session{AccessToken: accessToken, Expires: t}, nil
}

func fetchAccounts(loginToken string) (string, error) {
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
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	var payload models.AccountPayload
	// var objmap map[string]payload
	err = json.Unmarshal([]byte(data), &payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", payload)
	return "Hello", nil
}

func main() {
	db, err := models.InitDB()
	defer db.Close()

	creds := models.Login{}
	db.Last(&creds)

	if err != nil {
		panic(err)
	}
	userSession, err := getSession(creds.LoginToken)

	res, err := fetchAccounts(userSession.AccessToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
