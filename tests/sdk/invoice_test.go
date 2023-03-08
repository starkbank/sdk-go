package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestInvoicePost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	invoices, err := Invoice.Create(Example.Invoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}

	for _, invoice := range invoices {
		fmt.Printf("%+v", invoice)
	}

}

func TestInvoiceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var invoiceList []Invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoices := Invoice.Query(params, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	invoice, err := Invoice.Get(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, invoice.Id)
}

func TestInvoiceQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["after"] = "2020-04-01"
	params["before"] = "2020-04-30"

	invoices := Invoice.Query(params, nil)

	for invoice := range invoices {
		fmt.Printf("%+v", invoice)
	}
}

func TestInvoicePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	invoices, cursor, err := Invoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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

	var invoiceList []Invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoices := Invoice.Query(params, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	var patchData = map[string]interface{}{}
	patchData["amount"] = 49

	updated, err := Invoice.Update(invoiceList[rand.Intn(len(invoiceList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.Equal(t, updated.Amount, patchData["amount"])
}

func TestInvoiceQrcode(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var invoiceList []Invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoices := Invoice.Query(params, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	qrcode, err := Invoice.Qrcode(invoiceList[rand.Intn(len(invoiceList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.png", "invoice", invoiceList[rand.Intn(len(invoiceList))].Id)
	errFile := ioutil.WriteFile(filename, qrcode, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}

	assert.NotNil(t, qrcode)
}

func TestInvoicePdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var invoiceList []Invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoices := Invoice.Query(params, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	pdf, err := Invoice.Pdf(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "invoice", invoiceList[rand.Intn(len(invoiceList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}

func TestInvoicePayment(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var invoiceList []Invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(12)

	invoices := Invoice.Query(params, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	payment, err := Invoice.GetPayment(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Printf("%+v", payment)
	assert.NotNil(t, payment.Method)
}
