package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestBoletoPaymentPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, err := BoletoPayment.Create(Example.BoletosPayment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, payment := range payments {
		fmt.Printf("%+v", payment)
		assert.NotNil(t, payment.Id)
	}
}

func TestBoletoPaymentGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var paymentList []BoletoPayment.BoletoPayment
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	payments := BoletoPayment.Query(params, nil)
	for payment := range payments {
		paymentList = append(paymentList, payment)
	}

	payment, err := BoletoPayment.Get(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Printf("%+v", payment)
	assert.NotNil(t, payment.Id)
}

func TestBoletoPaymentPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var paymentList []BoletoPayment.BoletoPayment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "created"

	payments := BoletoPayment.Query(params, nil)
	for payment := range payments {
		paymentList = append(paymentList, payment)
	}

	pdf, err := BoletoPayment.Pdf(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "boleto-payment", paymentList[rand.Intn(len(paymentList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}

func TestBoletoPaymentQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 202

	payments := BoletoPayment.Query(params, nil)
	for payment := range payments {
		assert.NotNil(t, payment.Id)
		i++
	}
	assert.Equal(t, 202, i)
}

func TestBoletoPaymentPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BoletoPayment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, log := range logs {
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestBoletoPaymentPostDelete(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, errCreate := BoletoPayment.Create(Example.BoletosPayment(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}

	deleted, err := BoletoPayment.Delete(payments[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Printf("%+v", deleted)
	assert.NotNil(t, deleted.Id)
}
