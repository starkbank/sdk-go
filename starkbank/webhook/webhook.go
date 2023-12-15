package webhook

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

//	Webhook struct
//
//	A Webhook is used to subscribe to notification events on a user-selected endpoint.
//	Currently available services for subscription are transfer, boleto, boleto-holmes,
//	boleto-payment, brcode-payment, utility-payment, deposit and invoice.
//
//	Parameters (required):
//	- Url [string]: Url that will be notified when an event occurs.
//	- Subscriptions [slice of strings]: slice of any non-empty combination of the available services. ex: []string{"transfer", "invoice", "deposit"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the webhook is created. ex: "5656565656565656"

type Webhook struct {
	Url           string   `json:",omitempty"`
	Subscriptions []string `json:",omitempty"`
	Id            string   `json:",omitempty"`
}

var resource = map[string]string{"name": "Webhook"}

func Create(webhook Webhook, user user.User) (Webhook, Error.StarkErrors) {
	//	Create Webhook
	//
	//	Send a single Webhook for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- webhook [Webhook struct]: webhookData parameters for the creation of the webhook
	//		- url [string]: url to which notification events will be sent to. ex: "https://webhook.site/60e9c18e-4b5c-4369-bda1-ab5fcd8e1b29"
	//		- subscriptions [slice of strings]: slice of any non-empty combination of the available services. ex: ex: []string{"transfer", "boleto-payment"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Webhook struct with updated attributes
	var webHook Webhook
	create, err := utils.Single(resource, webhook, user)
	unmarshalError := json.Unmarshal(create, &webHook)
	if unmarshalError != nil {
		return webHook, err
	}
	return webHook, err
}

func Get(id string, user user.User) (Webhook, Error.StarkErrors) {
	//	Retrieve a specific Webhook by its id
	//
	//	Receive a single Webhook subscription struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Webhook struct with updated attributes
	var webhook Webhook
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &webhook)
	if unmarshalError != nil {
		return webhook, err
	}
	return webhook, err
}

func Query(params map[string]interface{}, user user.User) chan Webhook {
	//	Retrieve Webhook structs
	//
	//	Receive a channel of Webhook structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of Webhook structs with updated attributes
	var webhook Webhook
	webhooks := make(chan Webhook)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &webhook)
			if err != nil {
				panic(err)
			}
			webhooks <- webhook
		}
		close(webhooks)
	}()
	return webhooks
}

func Page(params map[string]interface{}, user user.User) ([]Webhook, string, Error.StarkErrors) {
	//	Retrieve paged Webhook structs
	//
	//	Receive a slice of up to 100 Webhook structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Webhook structs with updated attributes
	//	- cursor to retrieve the next page of Webhook structs
	var webhooks []Webhook
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &webhooks)
	if unmarshalError != nil {
		return webhooks, cursor, err
	}
	return webhooks, cursor, err
}

func Delete(id string, user user.User) (Webhook, Error.StarkErrors) {
	//	Delete a Webhook entity
	//
	//	Delete a Webhook entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: Webhook unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- deleted Webhook struct
	var webhook Webhook
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &webhook)
	if unmarshalError != nil {
		return webhook, err
	}
	return webhook, err
}
