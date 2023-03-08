package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	PaymentPreview "github.com/starkbank/sdk-go/starkbank/paymentpreview"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoletoPreviewPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	previews, err := PaymentPreview.Create(Example.PaymentPreviewBoleto(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, preview := range previews {
		fmt.Printf("%+v", preview)
		assert.NotNil(t, preview.Id)
		assert.Equal(t, preview.Type, "boleto-payment")
	}
}

func TestBrcodePreviewPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	previews, err := PaymentPreview.Create(Example.PaymentPreviewBrcode(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, preview := range previews {
		assert.NotNil(t, preview.Id)
		assert.Equal(t, preview.Type, "brcode-payment")
	}
}

func TestTaxPreviewPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	previews, err := PaymentPreview.Create(Example.PaymentPreviewTaxPreview(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, preview := range previews {
		assert.NotNil(t, preview.Id)
		assert.Equal(t, preview.Type, "tax-payment")
	}
}

func TestUtilityPreviewPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	previews, err := PaymentPreview.Create(Example.PaymentPreviewUtility(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, preview := range previews {
		assert.NotNil(t, preview.Id)
		assert.Equal(t, preview.Type, "utility-payment")
	}
}
