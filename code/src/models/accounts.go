package models

import (
	"github.com/jinzhu/gorm"
)

type AccountPayload struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	gorm.Model
	Id               string        `json:"id"`
	ProviderId       string        `json:"providerId"`
	Name             string        `json:"name"`
	Number           AccountNumber `json:"number"`
	BookedBalance    float32       `json:"bookedBalance"`
	AvailableBalance float32       `json:"availableBalance"`
	Currency         string        `json:"currency"`
	Type             string        `json:"type"`
	MigrationVersion int           `json:"migrationVersion"`
	IsPaymentAccount bool          `json:"isPaymentAccount"`
}

type AccountNumber struct {
	gorm.Model
	Bbantype   string      `json:"bbanType"`
	Bban       string      `json:"bban"`
	Iban       string      `json:"iban"`
	BbanParsed BankAccount `json:"bbanParsed"`
}

type BankAccount struct {
	gorm.Model
	Bankcode string `json:"bankCode"`
	Number   string `json:"accountNumber"`
}
