package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Deposit "github.com/starkbank/sdk-go/starkbank/deposit"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDepositGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var depositList []Deposit.Deposit
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	deposits := Deposit.Query(params, nil)
	for deposit := range deposits {
		depositList = append(depositList, deposit)
	}

	deposit, err := Deposit.Get(depositList[rand.Intn(len(depositList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	assert.NotNil(t, deposit)
}

func TestDepositQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	deposits := Deposit.Query(params, nil)

	for deposit := range deposits {
		assert.NotNil(t, deposit.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestDepositPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	deposits, cursor, err := Deposit.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, deposit := range deposits {
		ids = append(ids, deposit.Id)
		assert.NotNil(t, deposit.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
