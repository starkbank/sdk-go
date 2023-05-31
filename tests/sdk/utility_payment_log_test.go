package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	UtilityLog "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestUtilityLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []UtilityLog.Log
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	logs := UtilityLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	utility, err := UtilityLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}

	assert.NotNil(t, utility)
}

func TestUtilityLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := UtilityLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestUtilityLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := UtilityLog.Page(params, nil)
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
