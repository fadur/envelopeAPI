package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Login struct {
	gorm.Model
	Expires            time.Time `json:"expires"`
	Label              string    `json:"label"`
	LoginToken         string    `json:"loginToken"`
	ProviderId         string    `json:"providerId"`
	SubjectId          string    `json:"subjectId"`
	SupportsUnattended bool      `json:"supportsUnattended"`
}

type Session struct {
	gorm.Model
	AccessToken string    `json:accessToken`
	Expires     time.Time `json:expires`
}

type Payload struct {
	Login      Login
	ProviderId string `json:providerId`
	Session    Session
}

func (p *Payload) Save(db *gorm.DB) {
	db.Create(&p.Login)
	db.Create(&p.Session)
}
