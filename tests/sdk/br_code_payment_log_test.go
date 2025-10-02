package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	BrcodePaymentLog "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBrcodePaymentLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var brcodeList []BrcodePaymentLog.Log

	brcodes, errorChannel := BrcodePaymentLog.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	log, err := BrcodePaymentLog.Get(brcodeList[rand.Intn(len(brcodeList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, log)
}

func TestBrcodePaymentLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []BrcodePaymentLog.Log

	logs, errorChannel := BrcodePaymentLog.Query(params, nil)
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

func TestBrcodePaymentLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BrcodePaymentLog.Page(params, nil)
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
