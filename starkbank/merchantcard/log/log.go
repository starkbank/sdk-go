package log

import (
	"encoding/json"
	"time"
	MerchantCard "github.com/starkbank/sdk-go/starkbank/merchantcard"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

type Log struct {
	Id            string                    `json:",omitempty"`
	Card          MerchantCard.MerchantCard `json:",omitempty"`
	Errors        []interface{}             `json:",omitempty"`
	Type          string                    `json:",omitempty"`
	Created       *time.Time                `json:",omitempty"`
	TransactionId string                    `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantCardLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	var cardLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &cardLog)
	if unmarshalError != nil {
		return cardLog, err
	}
	return cardLog, err
}

func Query(params map[string]interface{}, user user.User) (chan Log, chan Error.StarkErrors) {
	var cardLog Log
	logs := make(chan Log)
	logsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &cardLog)
			if err != nil {
				logsError <- Error.UnknownError(err.Error())
				continue
			}
			logs <- cardLog
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
	var cardLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &cardLogs)
	if unmarshalError != nil {
		return cardLogs, cursor, err
	}
	return cardLogs, cursor, err
}
