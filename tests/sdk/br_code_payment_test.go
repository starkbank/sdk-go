package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBrcodePaymentPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, err := BrcodePayment.Create(Example.BrcodePayment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, payment := range payments {
		assert.NotNil(t, payment.Id)
	}
}

func TestBrcodePaymentGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var brcodeList []BrcodePayment.BrcodePayment

	brcodes, errorChannel := BrcodePayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	brcode, err := BrcodePayment.Get(brcodeList[rand.Intn(len(brcodeList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, brcode.Id)
}

func TestBrcodePaymentPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "success"
	
	var brcodeList []BrcodePayment.BrcodePayment

	brcodes, errorChannel := BrcodePayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	pdf, err := BrcodePayment.Pdf(brcodeList[rand.Intn(len(brcodeList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, pdf)
}

func TestBrcodePaymentQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var brcodeList []BrcodePayment.BrcodePayment

	brcodes, errorChannel := BrcodePayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	assert.Equal(t, limit, len(brcodeList))
}

func TestBrcodePaymentPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	brcodes, cursor, err := BrcodePayment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, brcode := range brcodes {
		ids = append(ids, brcode.Id)
		assert.NotNil(t, brcode.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestBrcodePaymentUpdate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "created"

	var brcodeList []BrcodePayment.BrcodePayment

	brcodes, errorChannel := BrcodePayment.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	if len(brcodeList) == 0 {
		t.Skip("No BrcodePayment with status created found")
	}

	var patchData = map[string]interface{}{}
	patchData["status"] = "canceled"

	updated, err := BrcodePayment.Update(brcodeList[rand.Intn(len(brcodeList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, updated)
}
