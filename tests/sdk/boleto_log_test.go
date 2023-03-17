package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	BoletoLog "github.com/starkbank/sdk-go/starkbank/boleto/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBoletoLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var boletoList []BoletoLog.Log
	var params = map[string]interface{}{}
	params["after"] = "2020-04-01"
	params["before"] = "2020-04-30"
	params["limit"] = 1

	boletos := BoletoLog.Query(params, nil)
	for boleto := range boletos {
		fmt.Printf("%+v", boleto)
		boletoList = append(boletoList, boleto)
	}

	log, err := BoletoLog.Get(boletoList[rand.Intn(len(boletoList))].Id, nil)
	if err.Errors != nil {
		fmt.Printf("%+v", log)
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
}

func TestBoletoLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 3

	logs := BoletoLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 3, i)
}

func TestBoletoLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BoletoLog.Page(params, nil)
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
