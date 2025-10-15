package log

import (
	"encoding/json"
	"time"
	MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

type Log struct {
	Id       string          				         `json:",omitempty"`
	Created  *time.Time      				         `json:",omitempty"`
	Type     string          				         `json:",omitempty"`
	Errors   []interface{}              		 `json:",omitempty"`
	Session  MerchantSession.MerchantSession `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantSessionLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	var log Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &log)
	if unmarshalError != nil {
		return log, err
	}
	return log, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	var log Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &log)
			if err != nil {
				panic(err)
			}
			logs <- log
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	var logs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &logs)
	if unmarshalError != nil {
		return logs, cursor, err
	}
	return logs, cursor, err
}
