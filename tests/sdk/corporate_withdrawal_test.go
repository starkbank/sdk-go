package sdk

import (
	"fmt"
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
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, withdrawal.Id)
	fmt.Println(withdrawal.Id)

}

func TestCorporateWithdrawalQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	withdrawals := CorporateWithdrawal.Query(params, nil)
	for withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
		fmt.Println(withdrawal.Id)
	}
}

func TestCorporateWithdrawalPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	withdrawals, cursor, err := CorporateWithdrawal.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
		fmt.Println(withdrawal.Id)
	}

	fmt.Println(cursor)
}

func TestCorporateWithdrawalGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var withdrawalList []CorporateWithdrawal.CorporateWithdrawal
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	withdrawals := CorporateWithdrawal.Query(paramsQuery, nil)
	for withdrawal := range withdrawals {
		withdrawalList = append(withdrawalList, withdrawal)
	}

	withdrawal, err := CorporateWithdrawal.Get(withdrawalList[rand.Intn(len(withdrawalList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, withdrawal.Id)
	fmt.Println(withdrawal.Id)
}
