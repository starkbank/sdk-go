package log

import (
	"encoding/json"
	"time"
	MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

type Log struct {
	Id            string                            `json:",omitempty"`
	Purchase      MerchantPurchase.MerchantPurchase `json:",omitempty"`
	Errors        []interface{}                     `json:",omitempty"`
	Type          string                            `json:",omitempty"`
	Created       *time.Time                        `json:",omitempty"`
	TransactionId string                            `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantPurchaseLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	var purchaseLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &purchaseLog)
	if unmarshalError != nil {
		return purchaseLog, err
	}
	return purchaseLog, err
}

func Query(params map[string]interface{}, user user.User) (chan Log, chan Error.StarkErrors) {
	var purchaseLog Log
	logs := make(chan Log)
	logsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &purchaseLog)
			if err != nil {
				logsError <- Error.UnknownError(err.Error())
				continue
			}
			logs <- purchaseLog
		}
		for err := range errorChannel {
			logsError <- err
		}
		close(logs)
		close(logsError)
	}()
	return logs, logsError
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	var purchaseLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &purchaseLogs)
	if unmarshalError != nil {
		return purchaseLogs, cursor, err
	}
	return purchaseLogs, cursor, err
}
