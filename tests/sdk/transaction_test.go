package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestTransactionPost(t *testing.T) {
	
	starkbank.User = Utils.ExampleProject

	_, err := Transaction.Create(Example.Transaction(), nil)
	for _, e := range err.Errors {
		assert.Equal(t, "Unknown exception encountered: Function deprecated since v1.2.0", e.Message)
	}
}

func TestTransactionGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var transactionList []Transaction.Transaction

	transactions, errorChannel := Transaction.Query(params, nil)
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

	transaction, err := Transaction.Get(transactionList[rand.Intn(len(transactionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, transaction.Id)
}

func TestTransactionQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var transactionList []Transaction.Transaction

	transactions, errorChannel := Transaction.Query(params, nil)
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

func TestTransactionPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	transactions, cursor, err := Transaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, transaction := range transactions {
		ids = append(ids, transaction.Id)
		assert.NotNil(t, transaction.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
