package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	MerchantInstallment "github.com/starkbank/sdk-go/starkbank/merchantinstallment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)


func TestMerchantInstallmentQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var installmentList []MerchantInstallment.MerchantInstallment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	installments := MerchantInstallment.Query(params, starkbank.User)
	for installment := range installments {
		installmentList = append(installmentList, installment)
	}

	installment, err := MerchantInstallment.Get(installmentList[rand.Intn(len(installmentList))].Id, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, installment.Id)
}

func TestMerchantInstallmentPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	installments, cursor, err := MerchantInstallment.Page(params, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, installments)
	assert.NotNil(t, cursor)
	assert.Greater(t, len(installments), 0)
	for _, installment := range installments {
		assert.NotNil(t, installment.Id)
	}
}
