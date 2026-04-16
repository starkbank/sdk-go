package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	SplitReceiverLog "github.com/starkbank/sdk-go/starkbank/splitreceiver/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitReceiverLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logIds, receiverIds []string

	logs, errorChannel := SplitReceiverLog.Query(params, nil)
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
		receiverLog, err := SplitReceiverLog.Get(ids, nil)
		if err.Errors != nil {
			for _, erro := range err.Errors {
				t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
			}
		}
		assert.NotNil(t, receiverLog)
		receiverIds = append(receiverIds, receiverLog.Id)
	}
	assert.Equal(t, logIds, receiverIds)
}

func TestSplitReceiverLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []SplitReceiverLog.Log

	logs, errorChannel := SplitReceiverLog.Query(params, nil)
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

func TestSplitReceiverLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := SplitReceiverLog.Page(params, nil)
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
