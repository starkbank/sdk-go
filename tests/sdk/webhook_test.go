package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Webhook "github.com/starkbank/sdk-go/starkbank/webhook"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestWebhookPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	webhook, err := Webhook.Create(Example.Webhook(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	fmt.Printf("%+v", webhook)
	assert.NotNil(t, webhook.Id)
}

func TestWebhookGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var webhookList []Webhook.Webhook
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	webhooks := Webhook.Query(params, nil)
	for webhook := range webhooks {
		webhookList = append(webhookList, webhook)
	}

	webhook, err := Webhook.Get(webhookList[rand.Intn(len(webhookList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, webhook.Id)
}

func TestWebhookQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 2

	webhooks := Webhook.Query(params, nil)

	for webhook := range webhooks {
		assert.NotNil(t, webhook.Id)
		i++
	}
	assert.Equal(t, params["limit"], i)
}

func TestWebhookPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	webhooks, cursor, err := Webhook.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, webhook := range webhooks {
		ids = append(ids, webhook.Id)
		assert.NotNil(t, webhook.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestWebhookDelete(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	webhook, errCreate := Webhook.Create(Example.Webhook(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	canceled, err := Webhook.Delete(webhook.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, canceled.Id)
}
