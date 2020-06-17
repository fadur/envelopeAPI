package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// "time"
)

type TransactionResponse struct {
	Response []Transactions `json:"transactions"`
}

func (res *TransactionResponse) Save(db *gorm.DB) {
	for _, t := range res.Response {
		fmt.Printf("%v <--\n", t.Category)
	}
}

type Category struct {
	ID    string `gorm:"primary_key" json:"id"`
	SetId string `json:"setId"`
}

type Transactions struct {
	ID           string  `gorm:"primary_key" json:"id"`
	Date         string  `json:"date"`
	CreationDate string  `json:"creationDate"`
	Text         string  `json:"text"`
	OriginalText string  `json:"originalText"`
	Amount       float32 `json:"amount"`
	ExtraData    string  `json:"extraData"`
	Type         string  `json:"type"`
	Currency     string  `json:"currency"`
	State        string  `json:"state"`
	CategoryId   uint
	Category     Category `gorm:"foreignkey:CategoryId"`
}
