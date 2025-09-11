package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	CorporateWithdrawal "github.com/starkbank/sdk-go/starkbank/corporatewithdrawal"
	"github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateWithdrawalPost(t *testing.T) {

	starkbank.User = utils.ExampleProject

	withdrawal, err := CorporateWithdrawal.Create(Example.CorporateWithdrawal(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, withdrawal.Id)
}

func TestCorporateWithdrawalQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var withdrawalList []CorporateWithdrawal.CorporateWithdrawal

	withdrawals, errorChannel := CorporateWithdrawal.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case withdrawal, ok := <-withdrawals:
			if !ok {
				break loop
			}
			withdrawalList = append(withdrawalList, withdrawal)
		}
	}

	assert.Equal(t, limit, len(withdrawalList))
}

func TestCorporateWithdrawalPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	withdrawals, _, err := CorporateWithdrawal.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
	}
}

func TestCorporateWithdrawalGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var withdrawalList []CorporateWithdrawal.CorporateWithdrawal

	withdrawals, errorChannel := CorporateWithdrawal.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case withdrawal, ok := <-withdrawals:
			if !ok {
				break loop
			}
			withdrawalList = append(withdrawalList, withdrawal)
		}
	}

	withdrawal, err := CorporateWithdrawal.Get(withdrawalList[rand.Intn(len(withdrawalList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, withdrawal.Id)
}
