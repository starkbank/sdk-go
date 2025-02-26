package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/merchantpurchase/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantPurchaseLogGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var logList []log.Log
	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := log.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	retrievedLog, err := log.Get(logList[0].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	if retrievedLog.Id != logList[0].Id {
		t.Fail()
	}
	fmt.Println(retrievedLog)
}

func TestMerchantPurchaseLogQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := log.Query(params, nil)

	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestMerchantPurchaseLogPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := log.Page(params, nil)
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
