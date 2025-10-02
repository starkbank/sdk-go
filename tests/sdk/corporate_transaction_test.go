package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	CorporateTransaction "github.com/starkbank/sdk-go/starkbank/corporatetransaction"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateTransactionQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var transactionList []CorporateTransaction.CorporateTransaction

	transactions, errorChannel := CorporateTransaction.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			transactionList = append(transactionList, transaction)
		}
	}

	assert.Equal(t, limit, len(transactionList))
}

func TestCorporateTransactionPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions, cursor, err := CorporateTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
	assert.NotNil(t, cursor)
}

func TestCorporateTransactionGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var transactionList []CorporateTransaction.CorporateTransaction

	transactions, errorChannel := CorporateTransaction.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			transactionList = append(transactionList, transaction)
		}
	}

	transaction, err := CorporateTransaction.Get(transactionList[rand.Intn(len(transactionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, transaction.Id)
}
