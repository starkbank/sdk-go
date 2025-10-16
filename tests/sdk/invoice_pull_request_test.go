package sdk

import (
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	Event "github.com/starkbank/sdk-go/starkbank/event"
	Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
	InvoicePullSubscription "github.com/starkbank/sdk-go/starkbank/invoicepullsubscription"
	InvoicePullRequest "github.com/starkbank/sdk-go/starkbank/invoicepullrequest"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
)

func TestInvoicePullRequestCreateAndCancel(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	invoices, err := Invoice.Create(Example.Invoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	invoice := invoices[0]

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription("push"), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	subscription := invoicePullSubscription[0]

	invoicePullRequest, err := InvoicePullRequest.Create(Example.InvoicePullRequest(invoice.Id, subscription.Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	request := invoicePullRequest[0]
	assert.NotNil(t, request.Id)
	assert.Equal(t, request.InvoiceId, invoice.Id)
	assert.Equal(t, request.SubscriptionId, subscription.Id)
	
	canceledRequest, err := InvoicePullRequest.Cancel(request.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.Equal(t, canceledRequest.Id, request.Id)
}

func TestInvoicePullRequestQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var requestsList []InvoicePullRequest.InvoicePullRequest

	requests, errorChannel := InvoicePullRequest.Query(params, nil)
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
			assert.NotNil(t, request.Id)
			requestsList = append(requestsList, request)
		}
	}

	assert.Equal(t, limit, len(requestsList))
}

func TestInvoicePullRequestQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, errorChannel := InvoicePullRequest.Query(params, nil)
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
			getRequest, err := InvoicePullRequest.Get(request.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, request.Id, getRequest.Id)
		}
	}
}

func TestInvoicePullRequestPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := InvoicePullRequest.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		for _, entity := range page {
			assert.False(t, ids[entity.Id])
			ids[entity.Id] = true
			count++
		}
		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}
	assert.Equal(t, 10, count)
}

func TestParseInvoicePullRequestEvent(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	content := "{\"event\": {\"created\": \"2025-07-25T17:36:41.040267+00:00\", \"id\": \"4805265536843776\", \"log\": {\"created\": \"2025-07-25T17:36:39.571648+00:00\", \"description\": \"\", \"errors\": [], \"id\": \"5789040171286528\", \"reason\": \"\", \"request\": {\"attemptType\": \"default\", \"created\": \"2025-07-25T17:36:37.201258+00:00\", \"displayDescription\": \"\", \"due\": \"2025-07-30T07:00:00+00:00\", \"externalId\": \"a15c4821d1c2413a82a4f3cfeee1315e\", \"id\": \"5397390693498880\", \"installmentId\": \"5424937942646784\", \"invoiceId\": \"5118508564217856\", \"status\": \"pending\", \"subscriptionId\": \"5181739848695808\", \"tags\": [], \"updated\": \"2025-07-25T17:36:39.571665+00:00\"}, \"type\": \"pending\"}, \"subscription\": \"invoice-pull-request\", \"workspaceId\": \"6235001133727744\"}}"
	validSignature := "MEUCIQCvbPc+mWLLL5nwvOBy/3MVJ3JU9fG/rNmyqmHtaeJA9wIgOR8Tw75MSj7lR9DPqhM62tlq+cFkbw14T4KmDBeC5rM="

	event, err := Event.Parse(content, validSignature, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, event)
}
