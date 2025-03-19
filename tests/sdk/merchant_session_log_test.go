package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	MerchantSessionLog "github.com/starkbank/sdk-go/starkbank/merchantsession/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantSessionLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []MerchantSessionLog.Log
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs := MerchantSessionLog.Query(params, nil)

	for log := range logs {
		logList = append(logList, log)
	}

	log, err := MerchantSessionLog.Get(logList[0].Id, nil)

	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	assert.NotNil(t, log)
}

func TestMerchantSessionLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 101

	logs := MerchantSessionLog.Query(params, nil)

	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 101, i)
}

func TestMerchantSessionLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := MerchantSessionLog.Page(params, nil)
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