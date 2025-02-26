package log

import (
	"encoding/json"
	"time"
	MerchantInstallment "github.com/starkbank/sdk-go/starkbank/merchantinstallment"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

type Log struct {
	Id       			string                    									`json:",omitempty"`
	Installment 	MerchantInstallment.MerchantInstallment  		`json:",omitempty"`
	Errors   			[]interface{}                  							`json:",omitempty"`
	Type     			string                    									`json:",omitempty"`
	Created  			*time.Time                									`json:",omitempty"`
	TransactionId string																			`json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantInstallmentLog"}


func Get(id string, user user.User) (Log, Error.StarkErrors) {
	var installmentLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &installmentLog)
	if unmarshalError != nil {
		return installmentLog, err
	}
	return installmentLog, err
}


func Query(params map[string]interface{}, user user.User) chan Log {
	var installmentLog Log
	installmentLogs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &installmentLog)
			if err != nil {
				panic(err)
			}
			installmentLogs <- installmentLog
		}
		close(installmentLogs)
	}()
	return installmentLogs
}


func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	var installmentLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &installmentLogs)
	if unmarshalError != nil {
		return installmentLogs, cursor, err
	}
	return installmentLogs, cursor, err
}
