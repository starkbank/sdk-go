package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporateHolderLog "github.com/starkbank/sdk-go/starkbank/corporateholder/log"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateHolderLogQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := CorporateHolderLog.Query(paramsQuery, nil)

	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestCorporateHolderLogPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := CorporateHolderLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}
	assert.NotNil(t, cursor)
}

func TestCorporateHolderLogGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var logList []CorporateHolderLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := CorporateHolderLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := CorporateHolderLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
}
