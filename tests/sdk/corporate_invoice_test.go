package sdk

import (
	"fmt"
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
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, invoice.Id)
}

func TestCorporateInvoiceQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices := CorporateInvoice.Query(params, nil)
	for invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}
}

func TestCorporateInvoicePage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices, cursor, err := CorporateInvoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}
	assert.NotNil(t, cursor)
}
