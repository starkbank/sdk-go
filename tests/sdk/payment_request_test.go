package sdk

import (
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, request := range requests {
		assert.NotNil(t, request)
		assert.NotNil(t, request.Payment)
	}
}

func TestPaymentRequestQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var requestList []PaymentRequest.PaymentRequest

	requests, errorChannel := PaymentRequest.Query("5763106043068416", params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			requestList = append(requestList, request)
		}
	}
	assert.Equal(t, limit, len(requestList))
}

func TestPaymentRequestPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	requests, cursor, err := PaymentRequest.Page("5763106043068416", nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, request := range requests {
		assert.NotNil(t, request)
		assert.NotNil(t, request.CenterId)
		assert.NotNil(t, cursor)
	}
}
