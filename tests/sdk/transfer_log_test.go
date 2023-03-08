package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	TransferLog "github.com/starkbank/sdk-go/starkbank/transfer/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestTransferLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []TransferLog.Log
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	logs := TransferLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	transfers, err := TransferLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	fmt.Printf("%+v", transfers)
}

func TestTransferLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := TransferLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestTransferLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := TransferLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, log := range logs {
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
