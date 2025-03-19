package merchantsession

import (
	"time"
)

type Purchase struct {
	Id                 string                 `json:",omitempty"`
	Amount             int                    
	InstallmentCount   int                    `json:",omitempty"`
	CardId             string                 `json:",omitempty"`
	CardExpiration     string                 `json:",omitempty"`
	CardNumber         string                 `json:",omitempty"`
	CardSecurityCode   string                 `json:",omitempty"`
	HolderName         string                 `json:",omitempty"`
	HolderEmail        string                 `json:",omitempty"`
	HolderPhone        string                 `json:",omitempty"`
	FundingType        string                 `json:",omitempty"`
	BillingCountryCode string                 `json:",omitempty"`
	BillingCity        string                 `json:",omitempty"`
	BillingStateCode   string                 `json:",omitempty"`
	BillingStreetLine1 string                 `json:",omitempty"`
	BillingStreetLine2 string                 `json:",omitempty"`
	BillingZipCode     string                 `json:",omitempty"`
	Metadata           map[string]interface{} `json:",omitempty"`
	CardEnding         string                 `json:",omitempty"`
	ChallengeMode      string                 `json:",omitempty"`
	ChallengeUrl       string                 `json:",omitempty"`
	Created            *time.Time             `json:",omitempty"`
	CurrencyCode       string                 `json:",omitempty"`
	EndToEndId         string                 `json:",omitempty"`
	Fee                int                    `json:",omitempty"`
	Network            string                 `json:",omitempty"`
	Source             string                 `json:",omitempty"`
	Status             string                 `json:",omitempty"`
	Tags               []string               `json:",omitempty"`
	Updated            *time.Time             `json:",omitempty"`
}

var purchase Purchase
var SubResourcePurchase = map[string]string{"name": "Purchase"}
