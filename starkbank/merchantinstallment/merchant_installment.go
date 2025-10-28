package merchantinstallment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// Check out our API Documentation at https://starkbank.com/docs/api#merchant-installment

type MerchantInstallment struct {
	Id             string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	Network        string     `json:",omitempty"`
	FundingType    string     `json:",omitempty"`
	PurchaseId     string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
	Due            *time.Time `json:",omitempty"`
	Fee            int        `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantInstallment"}

func Get(id string, user user.User) (MerchantInstallment, Error.StarkErrors){
	var merchantInstallment MerchantInstallment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantInstallment)
	if unmarshalError != nil {
		return merchantInstallment, err
	}
	return merchantInstallment, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantInstallment, chan Error.StarkErrors) {
	var merchantInstallment MerchantInstallment
	merchantInstallments := make(chan MerchantInstallment)
	merchantInstallmentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantInstallment)
			if err != nil {
				merchantInstallmentsError <- Error.UnknownError(err.Error())
				continue
			}
			merchantInstallments <- merchantInstallment
		}
		for err := range errorChannel {
			merchantInstallmentsError <- err
		}
		close(merchantInstallments)
		close(merchantInstallmentsError)
	}()
	return merchantInstallments, merchantInstallmentsError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantInstallment, string, Error.StarkErrors) {
	var merchantInstallments []MerchantInstallment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantInstallments)
	if unmarshalError != nil {
		return merchantInstallments, cursor, err
	}
	return merchantInstallments, cursor, err
}
