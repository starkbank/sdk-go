package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporateTransaction "github.com/starkbank/sdk-go/starkbank/corporatetransaction"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateTransactionQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions := CorporateTransaction.Query(params, nil)
	for transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
}

func TestCorporateTransactionPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions, cursor, err := CorporateTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
	assert.NotNil(t, cursor)
}

func TestCorporateTransactionGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var transactionList []CorporateTransaction.CorporateTransaction
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	transactions := CorporateTransaction.Query(paramsQuery, nil)
	for transaction := range transactions {
		transactionList = append(transactionList, transaction)
	}

	transaction, err := CorporateTransaction.Get(transactionList[rand.Intn(len(transactionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, transaction.Id)
}
