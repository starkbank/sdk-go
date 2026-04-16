package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	SplitProfileLog "github.com/starkbank/sdk-go/starkbank/splitprofile/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitProfileLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logIds, profileIds []string

	logs, errorChannel := SplitProfileLog.Query(params, nil)
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
			logIds = append(logIds, log.Id)
		}
	}

	for _, ids := range logIds {
		profileLog, err := SplitProfileLog.Get(ids, nil)
		if err.Errors != nil {
			for _, erro := range err.Errors {
				t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
			}
		}
		assert.NotNil(t, profileLog)
		profileIds = append(profileIds, profileLog.Id)
	}
	assert.Equal(t, logIds, profileIds)
}

func TestSplitProfileLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []SplitProfileLog.Log

	logs, errorChannel := SplitProfileLog.Query(params, nil)
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

func TestSplitProfileLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := SplitProfileLog.Page(params, nil)
	for _, log := range logs {
		if err.Errors != nil {
			for _, erro := range err.Errors {
				t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
			}
		}
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
