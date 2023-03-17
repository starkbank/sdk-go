package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestTaxPaymentPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, err := TaxPayment.Create(Example.TaxPayment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, payment := range payments {
		assert.NotNil(t, payment.Id)
	}
}

func TestTaxPaymentGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var paymentList []TaxPayment.TaxPayment
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	payments := TaxPayment.Query(params, nil)
	for payment := range payments {
		paymentList = append(paymentList, payment)
	}

	payment, err := TaxPayment.Get(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Printf("%+v", payment)
	assert.NotNil(t, payment.Id)
}

func TestTaxPaymentQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["status"] = "processing"
	params["limit"] = rand.Intn(100)

	payments := TaxPayment.Query(params, nil)

	for payment := range payments {
		assert.NotNil(t, payment.Id)
		i++
	}
}

func TestTaxPaymentPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var lines []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	payments, cursor, err := TaxPayment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, payment := range payments {
		lines = append(lines, payment.Id)
		assert.NotNil(t, payment.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, lines, 4)
}

func TestTaxPaymentCancel(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, err := TaxPayment.Create(Example.TaxPayment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	canceled, err := TaxPayment.Delete(payments[rand.Intn(len(payments))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, canceled.Id)
}

func TestTaxPaymentPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var paymentList []TaxPayment.TaxPayment
	var params = map[string]interface{}{}
	params["status"] = "success"

	payments := TaxPayment.Query(params, nil)
	for payment := range payments {
		paymentList = append(paymentList, payment)
	}

	pdf, err := TaxPayment.Pdf(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, pdf)
}
