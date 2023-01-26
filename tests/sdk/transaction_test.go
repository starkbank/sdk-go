package sdk

import (
	"fmt"
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

	transactions, err := Transaction.Create(Example.Transaction(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
}

func TestTransactionGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var transactionList []Transaction.Transaction
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	transactions := Transaction.Query(params, nil)
	for transaction := range transactions {
		transactionList = append(transactionList, transaction)
	}

	transaction, err := Transaction.Get(transactionList[rand.Intn(len(transactionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, transaction.Id)
}

func TestTransactionQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	transactions := Transaction.Query(params, nil)

	for transaction := range transactions {
		assert.NotNil(t, transaction.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestTransactionPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	transactions, cursor, err := Transaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, transaction := range transactions {
		ids = append(ids, transaction.Id)
		assert.NotNil(t, transaction.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
