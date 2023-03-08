package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	TaxPaymentLog "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestTaxPaymentLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []TaxPaymentLog.Log
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	logs := TaxPaymentLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	payment, err := TaxPaymentLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Printf("%+v", payment)
	assert.NotNil(t, payment.Id)
}

func TestTaxPaymentLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["types"] = "failed"
	params["limit"] = rand.Intn(100)

	payments := TaxPaymentLog.Query(params, nil)

	for payment := range payments {
		assert.Equal(t, payment.Type, "failed")
	}
}

func TestTaxPaymentLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	payments, cursor, err := TaxPaymentLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, payment := range payments {
		ids = append(ids, payment.Id)
		assert.NotNil(t, payment.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
