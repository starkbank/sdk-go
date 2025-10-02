package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"os"
	"math/rand"
	"testing"
)

func TestBoletoPaymentPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, err := BoletoPayment.Create(Example.BoletosPayment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, payment := range payments {
		assert.NotNil(t, payment.Id)
	}
}

func TestBoletoPaymentGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var paymentList []BoletoPayment.BoletoPayment

	payments, errorChannel := BoletoPayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case payment, ok := <-payments:
			if !ok {
				break loop
			}
			paymentList = append(paymentList, payment)
		}
	}
	
	payment, err := BoletoPayment.Get(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, payment.Id)
}

func TestBoletoPaymentPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "success"
	
	var paymentList []BoletoPayment.BoletoPayment

	payments, errorChannel := BoletoPayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case payment, ok := <-payments:
			if !ok {
				break loop
			}
			paymentList = append(paymentList, payment)
		}
	}
			
	pdf, err := BoletoPayment.Pdf(paymentList[rand.Intn(len(paymentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "boleto-payment", paymentList[rand.Intn(len(paymentList))].Id)
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		t.Errorf("error writing file: %v", errFile)
	}
	assert.NotNil(t, pdf)
}

func TestBoletoPaymentQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var paymentList []BoletoPayment.BoletoPayment

	payments, errorChannel := BoletoPayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case payment, ok := <-payments:
			if !ok {
				break loop
			}
			paymentList = append(paymentList, payment)
		}
	}
		
	assert.Equal(t, limit, len(paymentList))
}

func TestBoletoPaymentPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := BoletoPayment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
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
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}

	deleted, err := BoletoPayment.Delete(payments[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, deleted.Id)
}
