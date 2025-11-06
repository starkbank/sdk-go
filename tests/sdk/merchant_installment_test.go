package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantInstallment "github.com/starkbank/sdk-go/starkbank/merchantinstallment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMerchantInstallmentQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var installmentList []MerchantInstallment.MerchantInstallment

	installments, errorChannel := MerchantInstallment.Query(params, starkbank.User)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case installment, ok := <-installments:
			if !ok {
				break loop
			}
			installmentList = append(installmentList, installment)
		}
	}
	assert.Equal(t, limit, len(installmentList))

	installment, err := MerchantInstallment.Get(installmentList[0].Id, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, installment.Id)
}

func TestMerchantInstallmentPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 5

	installments, cursor, err := MerchantInstallment.Page(params, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, installments)
	assert.NotNil(t, cursor)
	assert.Greater(t, len(installments), 0)
	for _, installment := range installments {
		assert.NotNil(t, installment.Id)
	}
}
