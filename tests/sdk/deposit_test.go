package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	Deposit "github.com/starkbank/sdk-go/starkbank/deposit"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDepositGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var depositList []Deposit.Deposit

	deposits, errorChannel := Deposit.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case deposit, ok := <-deposits:
			if !ok {
				break loop
			}
			depositList = append(depositList, deposit)
		}
	}

	deposit, err := Deposit.Get(depositList[rand.Intn(len(depositList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, deposit)
}

func TestDepositQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var depositList []Deposit.Deposit

	deposits, errorChannel := Deposit.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case deposit, ok := <-deposits:
			if !ok {
				break loop
			}
			depositList = append(depositList, deposit)
		}
	}

	assert.Equal(t, limit, len(depositList))
}

func TestDepositPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	deposits, cursor, err := Deposit.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, deposit := range deposits {
		ids = append(ids, deposit.Id)
		assert.NotNil(t, deposit.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
