package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	CorporateInvoice "github.com/starkbank/sdk-go/starkbank/corporateinvoice"
	"github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorporateInvoicePost(t *testing.T) {

	starkbank.User = utils.ExampleProject

	invoice, err := CorporateInvoice.Create(Example.CorporateInvoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, invoice.Id)
}

func TestCorporateInvoiceQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var invoiceList []CorporateInvoice.CorporateInvoice

	invoices, errorChannel := CorporateInvoice.Query(params, nil)
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

func TestCorporateInvoicePage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices, cursor, err := CorporateInvoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}
	assert.NotNil(t, cursor)
}
