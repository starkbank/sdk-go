package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	CorporateCardLog "github.com/starkbank/sdk-go/starkbank/corporatecard/log"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateCardLogQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []CorporateCardLog.Log

	logs, errorChannel := CorporateCardLog.Query(params, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			logList = append(logList, log)
		}
	}

	assert.Equal(t, limit, len(logList))
}

func TestCorporateCardLogPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := CorporateCardLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}

	assert.NotNil(t, cursor)
}

func TestCorporateCardLogGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var logList []CorporateCardLog.Log

	logs, errorChannel := CorporateCardLog.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			logList = append(logList, log)
		}
	}

	log, err := CorporateCardLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, log.Id)
}
