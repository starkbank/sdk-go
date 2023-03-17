package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DepositLog "github.com/starkbank/sdk-go/starkbank/deposit/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logIds, depositsIds []string
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	logs := DepositLog.Query(params, nil)

	for log := range logs {
		fmt.Println(log.Id)
		logIds = append(logIds, log.Id)
	}
	for _, ids := range logIds {
		deposit, err := DepositLog.Get(ids, nil)
		if err.Errors != nil {
			for _, erro := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
			}
		}
		fmt.Printf("%+v", deposit)
		depositsIds = append(depositsIds, deposit.Id)
	}
	assert.Equal(t, logIds, depositsIds)
}

func TestDepositLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := DepositLog.Query(params, nil)

	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestDepositLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := DepositLog.Page(params, nil)
	for _, log := range logs {
		if err.Errors != nil {
			for _, erro := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
			}
		}
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
