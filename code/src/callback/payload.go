package main

import (
	"time"
)

type Login struct {
	Expires            time.Time `json:expires`
	Label              string    `json:label`
	LoginToken         string    `json:loginToken`
	ProviderId         string    `json:providerId`
	SubjectId          string    `json:subjectId`
	SupportsUnattended string    `json:supportsUnattended`
}

type Session struct {
	AccessToken string `json:accessToken`
	Expires     string `json:expires`
}

type Payload struct {
	Login      Login
	ProviderId string `json:providerId`
	Session    Session
}
