package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	PaymentRequest "github.com/starkbank/sdk-go/starkbank/paymentrequest"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentRequestPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	requests, err := PaymentRequest.Create(Example.PaymentRequest(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, request := range requests {
		assert.NotNil(t, request.Payment)
	}
}

func TestPaymentRequestQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "pending"
	params["limit"] = 10

	requests := PaymentRequest.Query("5763106043068416", params, nil)

	for request := range requests {
		fmt.Println(request)
		assert.Equal(t, request.Status, "pending")
	}
}

func TestPaymentRequestPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	requests, cursor, err := PaymentRequest.Page("5763106043068416", nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, request := range requests {
		fmt.Printf("%+v\n", request)
		assert.NotNil(t, request.CenterId)
		assert.NotNil(t, cursor)
	}
}
