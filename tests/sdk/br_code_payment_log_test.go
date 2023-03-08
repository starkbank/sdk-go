package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	BrcodePaymentLog "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBrcodePaymentLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var brcodeList []BrcodePaymentLog.Log
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	brcodes := BrcodePaymentLog.Query(params, nil)
	for brcode := range brcodes {
		brcodeList = append(brcodeList, brcode)
	}

	log, err := BrcodePaymentLog.Get(brcodeList[rand.Intn(len(brcodeList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	fmt.Printf("%+v", log)
	fmt.Println(log.Id)
}

func TestBrcodePaymentLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := BrcodePaymentLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestBrcodePaymentLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BrcodePaymentLog.Page(params, nil)
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
