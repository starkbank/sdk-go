package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporateCardLog "github.com/starkbank/sdk-go/starkbank/corporatecard/log"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateCardLogQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := CorporateCardLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestCorporateCardLogPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := CorporateCardLog.Page(params, nil)
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

func TestCorporateCardLogGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var logList []CorporateCardLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := CorporateCardLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := CorporateCardLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
}
