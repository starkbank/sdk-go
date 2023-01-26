package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DarfPaymentLog "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDarfPaymentLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []DarfPaymentLog.Log
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	logs := DarfPaymentLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := DarfPaymentLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	fmt.Println(log.Id)
}

func TestDarfPaymentLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := DarfPaymentLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestDarfPaymentLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := DarfPaymentLog.Page(params, nil)
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
