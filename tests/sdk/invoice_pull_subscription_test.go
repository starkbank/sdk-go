package sdk

import (
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	Event "github.com/starkbank/sdk-go/starkbank/event"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	InvoicePullSubscription "github.com/starkbank/sdk-go/starkbank/invoicepullsubscription"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoicePullSubscriptionPush(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "push"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	subscription := invoicePullSubscription[0]
	assert.NotNil(t, subscription.Id)
	assert.NotNil(t, subscription.Name)
	assert.Equal(t, subscriptionType, subscription.Type)
}

func TestCreateInvoicePullSubscriptionQrCode(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "qrcode"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	subscription := invoicePullSubscription[0]
	assert.NotNil(t, subscription.Id)
	assert.NotNil(t, subscription.Name)
	assert.NotNil(t, subscription.TaxId)
	assert.Equal(t, subscriptionType, subscription.Type)
}

func TestCreateInvoicePullSubscriptionPaymentAndOrQrCode(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "paymentAndOrQrcode"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	subscription := invoicePullSubscription[0]
	assert.NotNil(t, subscription.Id)
	assert.NotNil(t, subscription.Name)
	assert.NotNil(t, subscription.TaxId)
	assert.Equal(t, subscriptionType, subscription.Type)
}

func TestInvoicePullSubscriptionQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var subscriptionsList []InvoicePullSubscription.InvoicePullSubscription

	subscriptions, errorChannel := InvoicePullSubscription.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case subscription, ok := <-subscriptions:
			if !ok {
				break loop
			}
			assert.NotNil(t, subscription.Id)
			subscriptionsList = append(subscriptionsList, subscription)
		}
	}

	assert.Equal(t, limit, len(subscriptionsList))
}

func TestInvoicePullSubscriptionQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	subscriptions, errorChannel := InvoicePullSubscription.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case subscription, ok := <-subscriptions:
			if !ok {
				break loop
			}
			getSubscription, err := InvoicePullSubscription.Get(subscription.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, subscription.Id, getSubscription.Id)
		}
	}
}

func TestInvoicePullSubscriptionCreateAndCancel(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "push"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	subscription := invoicePullSubscription[0]

	// Change Subscription Status to "active"
	deleted, err := InvoicePullSubscription.Cancel(subscription.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.Equal(t, subscription.Id, deleted.Id)
}

func TestInvoicePullSubscriptionPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := InvoicePullSubscription.Page(params, nil)
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

func TestParseInvoicePullSubscriptionEvent(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	content := "{\"event\": {\"created\": \"2025-07-25T17:03:49.207194+00:00\", \"id\": \"5339088045473792\", \"log\": {\"created\": \"2025-07-25T17:03:47.305348+00:00\", \"description\": \"\", \"errors\": [], \"id\": \"4814349822590976\", \"reason\": \"\", \"subscription\": {\"amount\": 1500, \"amountMinLimit\": 0, \"bacenId\": \"RR3990842720250725fwJYIfdOGeF\", \"brcode\": \"00020101021226180014br.gov.bcb.pix5204000053039865802BR5925Stark Sociedade de Credit6009Sao Paulo62070503***80930014br.gov.bcb.pix2571brcode-h.sandbox.starkinfra.com/v2/rec/d2766b29d5184e90853405a9720439a16304686F\", \"created\": \"2025-07-25T17:03:47.280303+00:00\", \"data\": {}, \"displayDescription\": \"fist test - lucas4\", \"due\": \"2025-07-27T17:03:46.709858+00:00\", \"end\": \"2055-06-23T03:00:00+00:00\", \"externalId\": \"3581163bfe96436794a5284d5eb7a5b9\", \"id\": \"4786208928432128\", \"interval\": \"week\", \"name\": \"jaojao\", \"pullMode\": \"manual\", \"pullRetryLimit\": 3, \"referenceCode\": \"ricandalarrapateumanual\", \"start\": \"2055-06-16T03:00:00+00:00\", \"status\": \"created\", \"tags\": [], \"taxId\": \"457.965.518-41\", \"type\": \"qrcode\", \"updated\": \"2025-07-25T17:03:47.305366+00:00\"}, \"type\": \"created\"}, \"subscription\": \"invoice-pull-subscription\", \"workspaceId\": \"6235001133727744\"}}"
	validSignature := "MEYCIQDPoI8o1N1qbtq24wY2cvQ4reAxHv/AZs901L6WKU8ylwIhAJ4+ARRrgNARFu/1SbHnuDoHX4EtvkZDCmxTjP9WsT1b"

	event, err := Event.Parse(content, validSignature, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, event)
}
