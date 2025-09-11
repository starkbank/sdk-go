package sdk

import (
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, webhook.Id)
}

func TestWebhookGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var webhookList []Webhook.Webhook

	webhooks, errorChannel := Webhook.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case webhook, ok := <-webhooks:
			if !ok {
				break loop
			}
			webhookList = append(webhookList, webhook)
		}
	}

	webhook, err := Webhook.Get(webhookList[rand.Intn(len(webhookList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, webhook.Id)
}

func TestWebhookQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var webhookList []Webhook.Webhook

	webhooks, errorChannel := Webhook.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case webhook, ok := <-webhooks:
			if !ok {
				break loop
			}
			webhookList = append(webhookList, webhook)
		}
	}

	assert.Equal(t, limit, len(webhookList))
}

func TestWebhookPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	webhooks, cursor, err := Webhook.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
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
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	canceled, err := Webhook.Delete(webhook.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, canceled.Id)
}
