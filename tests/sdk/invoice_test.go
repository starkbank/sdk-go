package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"os"
	"math/rand"
	"testing"
)

func TestInvoicePost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	invoices, err := Invoice.Create(Example.Invoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}
}

func TestInvoiceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	invoice, err := Invoice.Get(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, invoice.Id)
}

func TestInvoiceQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	assert.Equal(t, limit, len(invoiceList))
}

func TestInvoicePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	invoices, cursor, err := Invoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, invoice := range invoices {
		ids = append(ids, invoice.Id)
		assert.NotNil(t, invoice.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestInvoiceUpdate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["amount"] = 49

	updated, err := Invoice.Update(invoiceList[rand.Intn(len(invoiceList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, updated.Amount, patchData["amount"])
}

func TestInvoiceQrcode(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	qrcode, err := Invoice.Qrcode(invoiceList[rand.Intn(len(invoiceList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.png", "invoice", invoiceList[rand.Intn(len(invoiceList))].Id)
	errFile := os.WriteFile(filename, qrcode, 0666)
	if errFile != nil {
		t.Errorf("error writing file: %s", errFile.Error())
	}

	assert.NotNil(t, qrcode)
}

func TestInvoicePdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	pdf, err := Invoice.Pdf(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "invoice", invoiceList[rand.Intn(len(invoiceList))].Id)
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		t.Errorf("error writing file: %s", errFile.Error())
	}
	assert.NotNil(t, pdf)
}

func TestInvoicePayment(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "paid"
	
	var invoiceList []Invoice.Invoice

	invoices, errorChannel := Invoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	payment, err := Invoice.GetPayment(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, payment.Name)
	assert.NotNil(t, payment.Method)
}
