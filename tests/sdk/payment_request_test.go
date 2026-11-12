package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
	PaymentRequest "github.com/starkbank/sdk-go/starkbank/paymentrequest"
	TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestPaymentRequestParseTaxAndDarfPayments(t *testing.T) {

	requests, err := PaymentRequest.ParseRequests([]PaymentRequest.PaymentRequest{
		{
			Type: "tax-payment",
			Payment: map[string]interface{}{
				"description": "Tax payment",
			},
		},
		{
			Type: "darf-payment",
			Payment: map[string]interface{}{
				"description": "Darf payment",
			},
		},
	})
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}
	assert.Len(t, requests, 2)
	assert.IsType(t, TaxPayment.TaxPayment{}, requests[0].Payment)
	assert.IsType(t, DarfPayment.DarfPayment{}, requests[1].Payment)
	assert.Equal(t, "Darf payment", requests[1].Payment.(DarfPayment.DarfPayment).Description)
}

func TestPaymentRequestParseDarfPaymentDates(t *testing.T) {

	for _, dates := range []map[string]interface{}{
		{"competence": "2022-10-28", "due": "2026-11-12", "scheduled": "2026-11-13"},
		{"competence": "2022-10-28T00:00:00+00:00", "due": "2026-11-12T00:00:00Z", "scheduled": "2026-11-13T00:00:00Z"},
	} {
		payment := map[string]interface{}{"description": "Darf payment"}
		for key, value := range dates {
			payment[key] = value
		}
		original := map[string]interface{}{}
		for key, value := range payment {
			original[key] = value
		}
		requests, err := PaymentRequest.ParseRequests([]PaymentRequest.PaymentRequest{
			{Type: "darf-payment", Payment: payment},
		})
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
			continue
		}
		darf := requests[0].Payment.(DarfPayment.DarfPayment)
		assert.Equal(t, 2022, darf.Competence.Year())
		assert.Equal(t, 12, darf.Due.Day())
		assert.Equal(t, 13, darf.Scheduled.Day())
		assert.Equal(t, original, payment, "the caller's map must not be mutated")
	}
}

func TestPaymentRequestParseDarfPaymentStructIsKept(t *testing.T) {

	due := time.Date(2026, 11, 12, 0, 0, 0, 0, time.UTC)
	darf := DarfPayment.DarfPayment{Description: "Darf payment", Due: &due}

	requests, err := PaymentRequest.ParseRequests([]PaymentRequest.PaymentRequest{
		{Type: "darf-payment", Payment: darf},
	})
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}
	assert.Equal(t, darf, requests[0].Payment)
}
