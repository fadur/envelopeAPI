package utils

import (
	"bytes"
	"encoding/json"
	models "envelopeApi/code/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"

	"io/ioutil"
	"time"
)

var client, secret string
const (

	base_url       = "https://api.nordicapigateway.com/v1/accounts"
	unattended_url = "https://api.nordicapigateway.com/v1/authentication/unattended"
	userHash       = "test-user-id"
)

func getSession(token string) (models.Session, error) {

	client	= os.Getenv("CLIENT")
	secret = os.Getenv("SECRET")
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
	expires := session["expires"].(string)
	t, err := time.Parse(time.RFC3339, expires)
	return models.Session{AccessToken: accessToken, Expires: t}, nil
}
func GetToken(db *gorm.DB) string {
	creds := models.Login{}
	db.Last(&creds)
	var err error
	// check if we have a session that hasn't expored
	var userSession = models.Session{}
	if db.Where("expires > ?", time.Now()).First(&userSession).RecordNotFound() {
		userSession, err = getSession(creds.LoginToken) // fetch a new session
		db.Create(&userSession)                         // save the current session
		if err != nil {
			panic(err)
		}

	}
	return userSession.AccessToken
}
