package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantSessionLog "github.com/starkbank/sdk-go/starkbank/merchantsession/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantSessionLogQueryAndGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, errorChannel := MerchantSessionLog.Query(params, nil)
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
			getLog, err := MerchantSessionLog.Get(log.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, log.Id, getLog.Id)
		}
	}
}

func TestMerchantSessionLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 101
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []MerchantSessionLog.Log

	logs, errorChannel := MerchantSessionLog.Query(params, nil)
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
			assert.NotNil(t, log.Id)
			logList = append(logList, log)
		}
	}
	assert.Equal(t, limit, len(logList))
}

func TestMerchantSessionLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := MerchantSessionLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, log := range logs {
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}

	assert.Len(t, ids, 4)
}
