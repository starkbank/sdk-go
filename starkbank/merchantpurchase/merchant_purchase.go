package merchantpurchase

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// Check out our API Documentation at https://starkbank.com/docs/api#merchant-purchase

type MerchantPurchase struct {
	Id                 string                 `json:",omitempty"`
	Amount             int                    `json:",omitempty"`
	InstallmentCount   int                    `json:",omitempty"`
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
	CardId             string                 `json:",omitempty"`
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

var resource = map[string]string{"name": "MerchantPurchase"}

func Create(merchantPurchase MerchantPurchase, user user.User) (MerchantPurchase, Error.StarkErrors) {
	create, err := utils.Single(resource, merchantPurchase, user)
	unmarshalError := json.Unmarshal(create, &merchantPurchase)
	if unmarshalError != nil {
		return merchantPurchase, err
	}
	return merchantPurchase, err
}

func Get(id string, user user.User) (MerchantPurchase, Error.StarkErrors) {
	var merchantPurchase MerchantPurchase
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantPurchase)
	if unmarshalError != nil {
		return merchantPurchase, err
	}
	return merchantPurchase, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantPurchase, chan Error.StarkErrors) {
	var merchantPurchase MerchantPurchase
	merchantPurchases := make(chan MerchantPurchase)
	merchantPurchasesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantPurchase)
			if err != nil {
				merchantPurchasesError <- Error.UnknownError(err.Error())
				continue
			}
			merchantPurchases <- merchantPurchase
		}
		for err := range errorChannel {
			merchantPurchasesError <- err
		}
		close(merchantPurchases)
		close(merchantPurchasesError)
	}()
	return merchantPurchases, merchantPurchasesError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantPurchase, string, Error.StarkErrors) {
	var merchantPurchases []MerchantPurchase
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantPurchases)
	if unmarshalError != nil {
		return merchantPurchases, cursor, err
	}
	return merchantPurchases, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (MerchantPurchase, Error.StarkErrors) {
	var purchase MerchantPurchase
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &purchase)
	if unmarshalError != nil {
		return purchase, err
	}
	return purchase, err
}
