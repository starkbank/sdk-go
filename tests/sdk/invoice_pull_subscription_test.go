package sdk

import (
	"fmt"
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	InvoicePullSubscription "github.com/starkbank/sdk-go/starkbank/invoicepullsubscription"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoicePullSubscriptionPush(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "push"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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

	var params = map[string]interface{}{}
	params["limit"] = 5

	subscriptions := InvoicePullSubscription.Query(params, nil)
	for subscription := range subscriptions {
		assert.NotNil(t, subscription.Id)
		assert.NotNil(t, subscription.Name)
	}
}

func TestInvoicePullSubscriptionQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	subscriptionsQuery := InvoicePullSubscription.Query(params, nil)
	firstSubscription := <-subscriptionsQuery
	subscription, err := InvoicePullSubscription.Get(firstSubscription.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, subscription.Id)
	assert.NotNil(t, subscription.Name)
}

func TestInvoicePullSubscriptionCreateAndCancel(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	subscriptionType := "push"

	invoicePullSubscription, err := InvoicePullSubscription.Create(Example.InvoicePullSubscription(subscriptionType), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	subscription := invoicePullSubscription[0]

	deleted, err := InvoicePullSubscription.Cancel(subscription.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.Equal(t, deleted.Id, subscription.Id)
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
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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