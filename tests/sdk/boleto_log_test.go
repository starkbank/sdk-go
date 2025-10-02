package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	BoletoLog "github.com/starkbank/sdk-go/starkbank/boleto/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBoletoLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var boletoList []BoletoLog.Log
	
	boletos, errorChannel := BoletoLog.Query(params, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case boleto, ok := <-boletos:
			if !ok {
				break loop
			}
			boletoList = append(boletoList, boleto)
		}
	}

	log, err := BoletoLog.Get(boletoList[rand.Intn(len(boletoList))].Id, nil)
	if err.Errors != nil {
		assert.NotNil(t, log)
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
}

func TestBoletoLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []BoletoLog.Log

	logs, errorChannel := BoletoLog.Query(params, nil)
	
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

func TestBoletoLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BoletoLog.Page(params, nil)
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
